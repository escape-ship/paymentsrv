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

// Amount represents payment amount information
type Amount struct {
	Total        int64 `json:"total"`
	TaxFree      int64 `json:"tax_free"`
	VAT          int64 `json:"vat"`
	Point        int64 `json:"point"`
	Discount     int64 `json:"discount"`
	GreenDeposit int64 `json:"green_deposit"`
}

// CardInfo represents card payment details
type CardInfo struct {
	KakaoPayPurchaseCorp     string `json:"kakaopay_purchase_corp"`
	KakaoPayPurchaseCorpCode string `json:"kakaopay_purchase_corp_code"`
	KakaoPayIssuerCorp       string `json:"kakaopay_issuer_corp"`
	KakaoPayIssuerCorpCode   string `json:"kakaopay_issuer_corp_code"`
	BIN                      string `json:"bin"`
	CardType                 string `json:"card_type"`
	InstallMonth             string `json:"install_month"`
	ApprovedID               string `json:"approved_id"`
	CardMID                  string `json:"card_mid"`
	InterestFreeInstall      string `json:"interest_free_install"`
	InstallmentType          string `json:"installment_type"`
	CardItemCode             string `json:"card_item_code"`
}

// ApproveResponse represents the response from the Approve API
type ApproveResponse struct {
	AID               string     `json:"aid"`
	TID               string     `json:"tid"`
	CID               string     `json:"cid"`
	SID               string     `json:"sid"`
	PartnerOrderID    string     `json:"partner_order_id"`
	PartnerUserID     string     `json:"partner_user_id"`
	PaymentMethodType string     `json:"payment_method_type"`
	Amount            Amount     `json:"amount"`
	CardInfo          *CardInfo  `json:"card_info,omitempty"`
	ItemName          string     `json:"item_name"`
	ItemCode          string     `json:"item_code"`
	Quantity          int32      `json:"quantity"`
	CreatedAt         time.Time  `json:"created_at"`
	ApprovedAt        time.Time  `json:"approved_at"`
	Payload           string     `json:"payload"`
}

// UnmarshalJSON custom unmarshal for datetime fields
func (a *ApproveResponse) UnmarshalJSON(data []byte) error {
	type Alias ApproveResponse
	aux := &struct {
		CreatedAt  string `json:"created_at"`
		ApprovedAt string `json:"approved_at"`
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Parse the created_at time format
	if aux.CreatedAt != "" {
		createdTime, err := time.Parse("2006-01-02T15:04:05", aux.CreatedAt)
		if err != nil {
			return fmt.Errorf("failed to parse created_at: %w", err)
		}
		a.CreatedAt = createdTime
	}

	// Parse the approved_at time format
	if aux.ApprovedAt != "" {
		approvedTime, err := time.Parse("2006-01-02T15:04:05", aux.ApprovedAt)
		if err != nil {
			return fmt.Errorf("failed to parse approved_at: %w", err)
		}
		a.ApprovedAt = approvedTime
	}

	return nil
}
