package handler

import (
	"net/http"

	"github.com/tuanta7/chasingpaper/internal/transport/rest/middleware"
	"github.com/tuanta7/chasingpaper/internal/usecase/plan"
	"github.com/tuanta7/chasingpaper/pkg/httpx"
)

type PlanHandler struct {
	uc *plan.UseCase
}

func NewPlanHandler(uc *plan.UseCase) *PlanHandler {
	return &PlanHandler{
		uc: uc,
	}
}

func (h *PlanHandler) ListPlans(w http.ResponseWriter, r *http.Request) {
	page, pageSize, _ := middleware.GetPaginationParams(r.Context())

	plans, err := h.uc.ListPlans(r.Context(), page, pageSize)
	if err != nil {
		_ = httpx.ErrorJSON(w, httpx.NewInternalError(httpx.WithHint(err.Error())))
		return
	}

	_ = httpx.ResponseJSON(w, http.StatusOK, httpx.JSON{
		"plans": plans,
	})
}

func (h *PlanHandler) GetPlanByID(w http.ResponseWriter, r *http.Request) {
}

func (h *PlanHandler) CreatePlan(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description" validate:"len=8"`
	}

	err := httpx.DecodeAndValidateJSON(r.Body, &requestPayload)
	if err != nil {
		_ = httpx.ErrorJSON(w, httpx.NewInvalidArgumentError(httpx.WithHint(err.Error())))
		return
	}
}

func (h *PlanHandler) UpdatePlan(w http.ResponseWriter, r *http.Request) {
}

func (h *PlanHandler) DeletePlan(w http.ResponseWriter, r *http.Request) {
}
