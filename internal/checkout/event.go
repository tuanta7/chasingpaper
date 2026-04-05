package checkout

import "time"

type WebhookEvent struct {
	ID        string    `json:"id"`
	Payload   []byte    `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
	ExpiresAt time.Time `json:"expires_at"`
}
