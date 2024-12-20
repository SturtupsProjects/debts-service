CREATE TABLE clients
(
    id                UUID PRIMARY KEY NOT NULL,
    full_name         VARCHAR NOT NULL,
    phone_number      VARCHAR(15) UNIQUE NOT NULL,
    address           VARCHAR NOT NULL,
    telegram_username VARCHAR DEFAULT NULL,
    telegram_user_id  INT UNIQUE,
    has_debt          BOOLEAN DEFAULT FALSE,
    client_status     VARCHAR(20) DEFAULT 'active' CHECK (client_status IN ('active', 'inactive', 'suspended')),
    notes             TEXT DEFAULT ' ',
    created_at        TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE installment
(
    id                UUID PRIMARY KEY,
    months_duration   INT NOT NULL CHECK (months_duration > 0),
    present_month     INT DEFAULT 1 CHECK (present_month > 0),
    client_id         UUID NOT NULL REFERENCES clients(id),
    total_amount      INT NOT NULL CHECK (total_amount >= 0),
    amount_paid       INT DEFAULT 0 CHECK (amount_paid >= 0),
    last_payment_date TIMESTAMP,
    is_fully_paid     BOOLEAN GENERATED ALWAYS AS (total_amount = amount_paid) STORED,
    created_at        TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_client FOREIGN KEY (client_id) REFERENCES clients (id) ON DELETE CASCADE
);

CREATE TABLE payments
(
    id             UUID PRIMARY KEY,
    installment_id UUID NOT NULL REFERENCES installment(id) ON DELETE CASCADE,
    payment_date   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    payment_amount INT NOT NULL CHECK (payment_amount > 0),
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_client_id ON installment (client_id);
CREATE INDEX idx_installment_id ON payments (installment_id);
