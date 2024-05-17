-- +goose Up
BEGIN;
DROP TABLE IF EXISTS order_deliveries CASCADE;
DROP TABLE IF EXISTS order_payments CASCADE;
DROP TABLE IF EXISTS order_items CASCADE;
DROP TABLE IF EXISTS orders CASCADE;

CREATE TYPE locale_type AS ENUM (
    'ru',
    'en'
);

CREATE TYPE currency_type AS ENUM (
    'RUB',
    'USD'
);

create table orders
(
    order_uid          VARCHAR(50) PRIMARY KEY,
    track_number       VARCHAR(100) NOT NULL,
    entry              VARCHAR(100) NOT NULL,
    locale             locale_type  NOT NULL,
    internal_signature VARCHAR(100) NOT NULL,
    customer_id        VARCHAR(100) NOT NULL,
    delivery_service   VARCHAR(100) NOT NULL,
    shardkey           VARCHAR(100) NOT NULL,
    sm_id              INTEGER      NOT NULL,
    date_created       TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    oof_shard          VARCHAR(100) NOT NULL
);

create table order_deliveries
(
    order_uid VARCHAR(50) PRIMARY KEY REFERENCES orders (order_uid),
    name      VARCHAR(100) NOT NULL,
    phone     VARCHAR(100) NOT NULL,
    zip       VARCHAR(100) NOT NULL,
    city      VARCHAR(100) NOT NULL,
    address   VARCHAR(100) NOT NULL,
    region    VARCHAR(100) NOT NULL,
    email     VARCHAR(100) NOT NULL
);

create table order_payments
(
    transaction   VARCHAR(50) PRIMARY KEY REFERENCES orders (order_uid),
    request_id    VARCHAR(100)  NOT NULL,
    currency      currency_type NOT NULL,
    provider      VARCHAR(100)  NOT NULL,
    amount        INTEGER       NOT NULL,
    payment_dt    BIGINT        NOT NULL,
    bank          VARCHAR(100)  NOT NULL,
    delivery_cost INTEGER       NOT NULL,
    goods_total   INTEGER       NOT NULL,
    custom_fee    INTEGER       NOT NULL
);

create table order_items
(
    id           SERIAL PRIMARY KEY,
    order_uid    VARCHAR(50) REFERENCES orders (order_uid),
    chrt_id      INTEGER      NOT NULL,
    track_number VARCHAR(100) NOT NULL,
    price        INTEGER      NOT NULL,
    rid          VARCHAR(100) NOT NULL,
    name         VARCHAR(100) NOT NULL,
    sale         INTEGER      NOT NULL,
    size         VARCHAR(100) NOT NULL,
    total_price  INTEGER      NOT NULL,
    nm_id        INTEGER      NOT NULL,
    brand        VARCHAR(100) NOT NULL,
    status       INTEGER      NOT NULL
);

COMMIT;

-- +goose Down
drop table order_deliveries;
drop table order_payments;
drop table order_items;
drop table orders;
