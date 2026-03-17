package kafka

import (
	"context"
	"errors"
	"sync"

	"github.com/IBM/sarama"
)

var (
	ErrProducerClosed = errors.New("kafka producer is closed")
	ErrEmptyTopic     = errors.New("kafka topic is required")
)

type SyncProducer struct {
	client   sarama.Client
	producer sarama.SyncProducer

	closeOnce sync.Once
	closeErr  error
}

type Message struct {
	Topic   string
	Key     []byte
	Value   []byte
	Headers map[string]string
}

func NewProducer(brokers []string, cfg *sarama.Config) (*SyncProducer, error) {
	if cfg == nil {
		cfg = sarama.NewConfig()
	}

	// Sync producers require successes to be enabled so callers can observe
	// the assigned partition and offset.
	cfg.Producer.Return.Successes = true

	client, err := sarama.NewClient(brokers, cfg)
	if err != nil {
		return nil, err
	}

	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		_ = client.Close()
		return nil, err
	}

	return &SyncProducer{
		client:   client,
		producer: producer,
	}, nil
}

func (p *SyncProducer) Send(ctx context.Context, message Message) (partition int32, offset int64, err error) {
	if p == nil || p.producer == nil {
		return 0, 0, ErrProducerClosed
	}

	if message.Topic == "" {
		return 0, 0, ErrEmptyTopic
	}

	if err := ctx.Err(); err != nil {
		return 0, 0, err
	}

	headers := make([]sarama.RecordHeader, 0, len(message.Headers))
	for key, value := range message.Headers {
		headers = append(headers, sarama.RecordHeader{
			Key:   []byte(key),
			Value: []byte(value),
		})
	}

	kafkaMessage := &sarama.ProducerMessage{
		Topic:   message.Topic,
		Key:     sarama.ByteEncoder(message.Key),
		Value:   sarama.ByteEncoder(message.Value),
		Headers: headers,
	}

	partition, offset, err = p.producer.SendMessage(kafkaMessage)
	if err != nil {
		return 0, 0, err
	}

	if err := ctx.Err(); err != nil {
		return partition, offset, err
	}

	return partition, offset, nil
}

func (p *SyncProducer) Close() error {
	if p == nil {
		return nil
	}

	p.closeOnce.Do(func() {
		if p.producer != nil {
			p.closeErr = p.producer.Close()
		}

		if p.client != nil {
			if err := p.client.Close(); err != nil && p.closeErr == nil {
				p.closeErr = err
			}
		}

		p.producer = nil
		p.client = nil
	})

	return p.closeErr
}
