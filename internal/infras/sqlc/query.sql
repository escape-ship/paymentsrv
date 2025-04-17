-- name: CreateRequest :one

INSERT INTO
    "payments".payment (
        transaction_id,
        status,
        order_id,
        user_id,
        item_name,
        item_quantity,
        total_amount,
        requested_at
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, NOW()) RETURNING *;

-- name: ApproveRequest :exec

UPDATE "payments".payment
SET
    approved_at = NOW()
WHERE transaction_id = $1;

-- name: GetRequestByTransactionID :one

SELECT
    transaction_id,
    status,
    order_id,
    user_id,
    item_name,
    item_quantity,
    total_amount,
    requested_at,
    approved_at
FROM "payments".payment
WHERE transaction_id = $1;

-- name: GetRequestByOrderID :one

SELECT
    transaction_id,
    status,
    order_id,
    user_id,
    item_name,
    item_quantity,
    total_amount,
    requested_at,
    approved_at
FROM "payments".payment
WHERE order_id = $1;