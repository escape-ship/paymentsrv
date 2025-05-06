START TRANSACTION;

CREATE SCHEMA IF NOT EXISTS paymentsrv;

CREATE TABLE paymentsrv.kakao (
    tid CHAR(20) PRIMARY KEY NOT NULL,
    status VARCHAR(24) NOT NULL,
    aid CHAR(20) NULL,
    partner_order_id CHAR(24) NOT NULL,
    partner_user_id CHAR(24) NOT NULL,
    item_name VARCHAR(255) NOT NULL,
    quantity INT NOT NULL,
    total_amount BIGINT NOT NULL,
    tax_free_amount BIGINT NOT NULL,
    created_at TIMESTAMP NULL,
    update_at TIMESTAMP NULL,
    approved_at TIMESTAMP NULL
);

COMMIT;