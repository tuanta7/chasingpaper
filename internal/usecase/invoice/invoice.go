package invoice

type Invoice struct {
	ID         string `json:"id"`
	ExternalID string `json:"external_id"`
	Provider   string `json:"provider"`
	Status     string `json:"status"`
	CreatedAt  string `json:"created_at"`
}
