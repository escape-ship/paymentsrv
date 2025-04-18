// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package postgresql

import (
	"database/sql"
)

type PaymentsPayment struct {
	TransactionID string       `json:"transaction_id"`
	Status        string       `json:"status"`
	OrderID       string       `json:"order_id"`
	UserID        string       `json:"user_id"`
	ItemName      string       `json:"item_name"`
	ItemQuantity  int32        `json:"item_quantity"`
	TotalAmount   int32        `json:"total_amount"`
	RequestedAt   sql.NullTime `json:"requested_at"`
	ApprovedAt    sql.NullTime `json:"approved_at"`
}
