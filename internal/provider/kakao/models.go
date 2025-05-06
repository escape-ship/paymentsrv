package kakao

import (
	"encoding/json"
	"fmt"
	"time"
)

// ReadyRequest represents the request parameters for the Ready API
type ReadyRequest struct {
	PartnerOrderID string
	PartnerUserID  string
	ItemName       string
	Quantity       int32
	TotalAmount    int64
	TaxFreeAmount  int64
	ApprovalURL    string
	FailURL        string
	CancelURL      string
}

// ReadyResponse represents the response from the Ready API
type ReadyResponse struct {
	TID                   string `json:"tid"`
	NextRedirectAppURL    string `json:"next_redirect_app_url"`
	NextRedirectMobileURL string `json:"next_redirect_mobile_url"`
	NextRedirectPCURL     string `json:"next_redirect_pc_url"`
	AndroidAppScheme      string `json:"android_app_scheme"`
	IOSAppScheme          string `json:"ios_app_scheme"`
}

// ApproveRequest represents the request parameters for the Approve API
type ApproveRequest struct {
	TID            string
	PartnerOrderID string
	PartnerUserID  string
	PgToken        string
}

// ApproveResponse represents the response from the Approve API
type ApproveResponse struct {
	AID            string    `json:"aid"`
	TID            string    `json:"tid"`
	PartnerOrderID string    `json:"partner_order_id"`
	PartnerUserID  string    `json:"partner_user_id"`
	ApprovedAt     time.Time `json:"approved_at"`
}

// UnmarshalJSON custom unmarshal for ApprovedAt
func (a *ApproveResponse) UnmarshalJSON(data []byte) error {
	type Alias ApproveResponse
	aux := &struct {
		ApprovedAt string `json:"approved_at"`
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Parse the custom time format
	parsedTime, err := time.Parse("2006-01-02T15:04:05", aux.ApprovedAt)
	if err != nil {
		return fmt.Errorf("failed to parse approved_at: %w", err)
	}
	a.ApprovedAt = parsedTime
	return nil
}
