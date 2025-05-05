START TRANSACTION;

CREATE SCHEMA IF NOT EXISTS payments;

CREATE TABLE payments.payment (
    transaction_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    status VARCHAR(24) NOT NULL,
    order_id VARCHAR(24) NOT NULL,
    user_id VARCHAR(24) NOT NULL,
    item_name VARCHAR(255) NOT NULL,
    item_quantity INT NOT NULL,
    total_amount INT NOT NULL,
    requested_at TIMESTAMP NULL,
    approved_at TIMESTAMP NULL
);

COMMIT;