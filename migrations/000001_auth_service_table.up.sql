-- Table: clients
CREATE TABLE clients
(
    id                UUID PRIMARY KEY NOT NULL,
    full_name         VARCHAR NOT NULL,
    phone_number      VARCHAR(13) UNIQUE NOT NULL ,
    address           VARCHAR NOT NULL ,
    telegram_username VARCHAR DEFAULT ' ',
    telegram_user_id  INT UNIQUE,
    has_debt          BOOLEAN   DEFAULT FALSE,
    client_status     VARCHAR   DEFAULT 'active',
    notes             TEXT      DEFAULT 'no notes',
    created_at        TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table: installment
CREATE TABLE installment
(
    id                UUID PRIMARY KEY,
    months_duration   INT NOT NULL,
    present_month     INT       DEFAULT 1 id UUID NOT NULL REFERENCES clients(client_id),
    total_amount      INT NOT NULL,
    amount_paid       INT       DEFAULT 0,
    last_payment_date TIMESTAMP,
    is_fully_paid     BOOLEAN GENERATED ALWAYS AS (total_amount = amount_paid) STORED,
    created_at        TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table: payments
CREATE TABLE payments
(
    id             UUID PRIMARY KEY,
    installment_id UUID NOT NULL REFERENCES installment (id),
    payment_date   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    payment_amount INT  NOT NULL,
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
