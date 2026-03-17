package main

import (
	"context"
	"log"
	"os"

	"github.com/tuanta7/chasingpaper/internal/config"
	"github.com/tuanta7/chasingpaper/internal/plan"
	"github.com/tuanta7/chasingpaper/internal/transport/rest"
	"github.com/tuanta7/chasingpaper/internal/transport/rest/handlers"
	"github.com/tuanta7/chasingpaper/pkg/graceful"
	"github.com/tuanta7/chasingpaper/pkg/monitor"
	"github.com/tuanta7/chasingpaper/pkg/postgres"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Commands: []*cli.Command{},
		Action: func(ctx context.Context, command *cli.Command) error {
			cfg := config.LoadConfig()
			initMonitor(ctx, cfg)

			pool, err := postgres.NewInstrumentedPool(ctx, cfg.Postgres.DSN)
			if err != nil {
				return err
			}

			planRepo := plan.NewPostgresRepository(pool)
			planUC := plan.NewUseCase(planRepo)
			planHandler := handlers.NewPlanHandler(planUC)

			server := rest.NewServer(cfg.BindAddress, planHandler)
			return graceful.StartServerWithGracefulShutdown(server)
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func initMonitor(ctx context.Context, cfg *config.Config) {
	if !cfg.EnableMetrics {
		_ = monitor.InitNoopMeterProvider()
	} else {
		_, err := monitor.InitMeterProvider(ctx, cfg.ServiceName, nil)
		if err != nil {
			log.Fatalf("Failed to initialize meter provider: %v", err)
		}
	}

	if !cfg.EnableTracing {
		_ = monitor.InitNoopTracerProvider()
	} else {
		_, err := monitor.InitTracerProvider(ctx, cfg.ServiceName, nil)
		if err != nil {
			log.Fatalf("Failed to initialize tracer provider: %v", err)
		}
	}
}
