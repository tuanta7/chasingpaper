package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.39.0"
	"go.opentelemetry.io/otel/trace"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type InstrumentedPool struct {
	pool   *pgxpool.Pool
	tracer trace.Tracer
	meter  metric.Meter
}

func NewInstrumentedPool(ctx context.Context, dsn string) (*InstrumentedPool, error) {
	pgxPool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	if pgxPool.Ping(ctx) != nil {
		return nil, err
	}

	p := &InstrumentedPool{
		pool:   pgxPool,
		tracer: otel.Tracer("postgres_tracer"),
		meter:  otel.Meter("postgres_meter"),
	}

	err = initMetrics(p.meter)
	return p, err
}

func (p *InstrumentedPool) Close() {
	p.pool.Close()
}

func (p *InstrumentedPool) Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error) {
	ctx, span := p.tracer.Start(ctx, "postgres_exec", trace.WithAttributes(
		semconv.DBSystemNamePostgreSQL,
		semconv.DBQueryText(sql),
	))
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()

	return p.pool.Exec(ctx, sql, arguments...)
}

func (p *InstrumentedPool) Query(ctx context.Context, sql string, args ...any) (rows pgx.Rows, err error) {
	start := time.Now()
	ctx, span := p.tracer.Start(ctx, "postgres_query", trace.WithAttributes(
		semconv.DBSystemNamePostgreSQL,
		semconv.DBQueryText(sql),
	))
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
		queryDuration.Record(ctx, time.Since(start).Seconds(), metric.WithAttributes(
			attribute.String("db.operation", "query"),
		))
	}()

	return p.pool.Query(ctx, sql, args...)
}

func (p *InstrumentedPool) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	start := time.Now()
	ctx, span := p.tracer.Start(ctx, "postgres_query_row", trace.WithAttributes(
		semconv.DBSystemNamePostgreSQL,
		semconv.DBQueryText(sql),
	))
	defer func() {
		span.End()
		queryDuration.Record(ctx, time.Since(start).Seconds(), metric.WithAttributes(
			attribute.String("db.operation", "query_row"),
		))
	}()

	return p.pool.QueryRow(ctx, sql, args...)
}

type TxStatement func(tx pgx.Tx) error

func (p *InstrumentedPool) DoTx(ctx context.Context, txOptions pgx.TxOptions, statements []TxStatement) (err error) {
	start := time.Now()
	ctx, span := p.tracer.Start(ctx, "postgres_do_tx", trace.WithAttributes(
		semconv.DBSystemNamePostgreSQL,
	))
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
		queryDuration.Record(ctx, time.Since(start).Seconds(), metric.WithAttributes(
			attribute.String("db.operation", "do_tx"),
		))
	}()

	tx, err := p.pool.BeginTx(ctx, txOptions)
	if err != nil {
		return err
	}

	for _, s := range statements {
		if se := s(tx); se != nil {
			_ = tx.Rollback(ctx)
			return err
		}
	}

	return tx.Commit(ctx)
}
