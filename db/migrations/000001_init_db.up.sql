CREATE TABLE orders
(
    order_uid          varchar(255) PRIMARY KEY,
    track_number       varchar(255) NOT NULL,
    entry              varchar(255) NOT NULL,
    locale             varchar(255) NOT NULL,
    internal_signature varchar(255) NOT NULL,
    customer_id        varchar(255) NOT NULL,
    delivery_service   varchar(255) NOT NULL,
    sm_id              integer      NOT NULL,
    date_created       timestamp(0) NOT NULL
);

CREATE TABLE deliveries
(
    delivery_uid uuid PRIMARY KEY,
    order_uid    varchar(255) NOT NULL REFERENCES orders (order_uid),
    name         varchar(255) NOT NULL,
    phone        varchar(255) NOT NULL,
    zip          varchar(255) NOT NULL,
    city         varchar(255) NOT NULL,
    address      varchar(255) NOT NULL,
    region       varchar(255) NOT NULL,
    email        varchar(255) NOT NULL
);

CREATE TABLE payments
(
    payment_uid   uuid PRIMARY KEY,
    order_uid     varchar(255) NOT NULL REFERENCES orders (order_uid),
    transaction   varchar(255) NOT NULL,
    request_id    varchar(255) NOT NULL,
    currency      varchar(255) NOT NULL,
    provider      varchar(255) NOT NULL,
    amount        integer      NOT NULL,
    payment_dt    timestamp(0) NOT NULL,
    bank          varchar(255) NOT NULL,
    delivery_cost integer      NOT NULL CHECK ( delivery_cost >= 0 ),
    goods_total   integer      NOT NULL CHECK ( goods_total > 0 ),
    custom_fee    integer      NOT NULL CHECK ( custom_fee >= 0 ),
    CHECK ( amount = delivery_cost + goods_total + custom_fee )
);

CREATE TABLE items
(
    item_uid     uuid PRIMARY KEY,
    order_uid    varchar(255) NOT NULL REFERENCES orders (order_uid),
    chrt_id      integer      NOT NULL,
    track_number varchar(255) NOT NULL,
    price        integer      NOT NULL CHECK ( price > 0 ),
    rid          varchar(255) NOT NULL,
    name         varchar(255) NOT NULL,
    sale         integer      NOT NULL CHECK ( sale >= 0 AND sale < 100 ),
    size         varchar(255) NOT NULL,
    total_price  integer      NOT NULL,
    nm_id        integer      NOT NULL,
    brand        varchar(255) NOT NULL,
    status       integer      NOT NULL
);

