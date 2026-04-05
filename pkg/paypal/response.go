package paypal

type link struct {
	Href    string `json:"href"`
	Rel     string `json:"rel"`
	Method  string `json:"method"`
	EncType string `json:"encType,omitempty"`
}

type errorResponse struct {
	Name    string         `json:"name"`
	Message string         `json:"message"`
	DebugID string         `json:"debug_id"`
	Details []errorDetails `json:"details"`
	Links   []link         `json:"links"`
}

type errorDetails struct {
	Issue       string `json:"issue"`
	Description string `json:"description"`
}
