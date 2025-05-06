-- name: CreateKakao :one

INSERT INTO
    "paymentsrv".kakao (
        tid,
        status,
        partner_order_id,
        partner_user_id,
        item_name,
        quantity,
        total_amount,
        tax_free_amount,
        created_at,
        update_at
    )
VALUES ($1, 'READY', $2, $3, $4, $5, $6, $7, NOW(), NOW()) RETURNING tid;

-- name: UpdateKakaoStatus :exec

UPDATE "paymentsrv".kakao
SET
    status = $1,
    update_at = NOW()
WHERE tid = $2;

-- name: UpdateKakaoApprove :exec

UPDATE "paymentsrv".kakao
SET
    status = 'APPROVED',
    aid = $2,
    approved_at = $3,
    update_at = NOW()
WHERE tid = $1;

-- name: GetKakaoByTID :one

SELECT
    *
FROM "paymentsrv".kakao
WHERE tid = $1;

-- name: GetKakaoByOrderID :one

SELECT
    *
FROM "paymentsrv".kakao
WHERE partner_order_id = $1;