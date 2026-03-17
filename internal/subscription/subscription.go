package subscription

type status string

const (
	StatusTrialing status = "trialing"
	StatusActive   status = "active"
	StatusExpired  status = "expired"
	StatusCanceled status = "canceled"
)

type Subscription struct {
	UserID      string `json:"user_id"`
	PlanID      string `json:"plan_id"`
	Status      status `json:"status"`
	AutoRenew   bool   `json:"auto_renew"`
	PeriodStart string `json:"period_start"`
	PeriodEnd   string `json:"period_end"`
}
