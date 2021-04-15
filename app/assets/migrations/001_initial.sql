-- +migrate Up
CREATE TABLE blocks
(
    id     BIGINT NOT NULL PRIMARY KEY,
    hash   text   NOT NULL,
    parent text   NOT NULL
);

CREATE TABLE transactions
(
    hash      TEXT        NOT NULL PRIMARY KEY,
    block_id  BIGINT      NOT NULL REFERENCES blocks (id) ON DELETE CASCADE,
    sender    text        NOT NULL,
    recipient text,
    value     NUMERIC(32) NOT NULL,
    nonce     BIGINT      NOT NULL,
    gas_price NUMERIC(32) NOT NULL,
    gas       BIGINT      NOT NULL,
    input     bytea,
    timestamp BIGINT      NOT NULL
);

CREATE INDEX tx_sender_idx ON transactions (sender);
CREATE INDEX tx_recipient_idx ON transactions (recipient);

CREATE TABLE payments
(
    tx_hash   TEXT        NOT NULL PRIMARY KEY REFERENCES transactions (hash) ON DELETE CASCADE,
    sender    text        NOT NULL,
    recipient text        NOT NULL,
    token     text,
    value     NUMERIC(32) NOT NULL,
    timestamp BIGINT      NOT NULL
);

CREATE INDEX pay_sender_idx ON payments(sender);
CREATE INDEX pay_recipient_idx ON payments(recipient);

-- +migrate Down
DROP TABLE payments;
DROP TABLE transactions;
DROP TABLE blocks;
