package point

type Wallet struct {
	UserID  string `json:"user_id"`
	Balance int64  `json:"balance"`
}

type WalletAudit struct {
	UserID    string `json:"user_id"`
	EventType string `json:"event_type"`
	Amount    int64  `json:"amount"`
	CreatedAt string `json:"created_at"`
}
