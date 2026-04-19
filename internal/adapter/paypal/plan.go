package paypal

import (
	"context"
	"encoding/json"
	"net/http"
)

type Plan struct {
	ID          string `json:"id"`
	Version     int64  `json:"version"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	Description string `json:"description"`
	UsageType   string `json:"usage_type"`
	CreateTime  string `json:"create_time"`
	Links       []Link `json:"links"`
}

type PlanDetails struct {
	Plan Plan `json:"plan"`
}

type listPlansResponse struct {
	TotalItems int    `json:"total_items"`
	TotalPages int    `json:"total_pages"`
	Plans      []Plan `json:"plans"`
	Links      []Link `json:"links"`
}

func (c *Client) ListPlans(ctx context.Context) ([]Plan, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+pathV2Plans, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data listPlansResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Plans, nil
}

func (c *Client) GetPlan(ctx context.Context, planID string) (*Plan, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+pathV2Plans+planID, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return nil, nil
}
