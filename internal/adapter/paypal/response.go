package paypal

type Link struct {
	Href    string `json:"href"`
	Rel     string `json:"rel"`
	Method  string `json:"method"`
	EncType string `json:"encType,omitempty"`
}

type ErrorResponse struct {
	Name    string         `json:"name"`
	Message string         `json:"message"`
	DebugID string         `json:"debug_id"`
	Details []ErrorDetails `json:"details"`
	Links   []Link         `json:"links"`
}

type ErrorDetails struct {
	Issue       string `json:"issue"`
	Description string `json:"description"`
}
