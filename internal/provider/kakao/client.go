package kakao

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Config represents the configuration for the Kakao Pay client
type Config struct {
	BaseURL   string
	SecretKey string
	CID       string
}

// Client represents a Kakao Pay API client
type Client struct {
	config     Config
	httpClient *http.Client
}

// NewClient creates a new Kakao Pay client
func NewClient(config Config) *Client {
	return &Client{
		config:     config,
		httpClient: &http.Client{},
	}
}

// Ready initiates a payment process
func (c *Client) Ready(ctx context.Context, req ReadyRequest) (*ReadyResponse, error) {
	reqMap := map[string]any{
		"cid":              c.config.CID,
		"partner_order_id": req.PartnerOrderID,
		"partner_user_id":  req.PartnerUserID,
		"item_name":        req.ItemName,
		"quantity":         req.Quantity,
		"total_amount":     req.TotalAmount,
		"tax_free_amount":  req.TaxFreeAmount,
		"approval_url":     req.ApprovalURL,
		"fail_url":         req.FailURL,
		"cancel_url":       req.CancelURL,
	}

	resp, err := c.doRequest(ctx, "ready", reqMap)
	if err != nil {
		return nil, fmt.Errorf("failed to make ready request: %w", err)
	}

	var readyResp ReadyResponse
	if err := json.Unmarshal(resp, &readyResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal ready response: %w", err)
	}

	return &readyResp, nil
}

// Approve approves a payment process
func (c *Client) Approve(ctx context.Context, req ApproveRequest) (*ApproveResponse, error) {
	reqMap := map[string]any{
		"cid":              c.config.CID,
		"tid":              req.TID,
		"partner_order_id": req.PartnerOrderID,
		"partner_user_id":  req.PartnerUserID,
		"pg_token":         req.PgToken,
	}

	resp, err := c.doRequest(ctx, "approve", reqMap)
	if err != nil {
		return nil, fmt.Errorf("failed to make approve request: %w", err)
	}

	var approveResp ApproveResponse
	if err := json.Unmarshal(resp, &approveResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal approve response: %w", err)
	}

	return &approveResp, nil
}

func (c *Client) doRequest(ctx context.Context, endpoint string, payload interface{}) ([]byte, error) {
	reqBody, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	url := fmt.Sprintf("%s/%s", c.config.BaseURL, endpoint)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "SECRET_KEY "+c.config.SecretKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	return body, nil
}
