CREATE TABLE clients
(
    id                UUID PRIMARY KEY   NOT NULL,
    full_name         VARCHAR            NOT NULL,
    phone_number      VARCHAR(15) UNIQUE NOT NULL,
    address           VARCHAR            NOT NULL,
    telegram_username VARCHAR     DEFAULT NULL,
    telegram_user_id  INT UNIQUE,
    has_debt          BOOLEAN     DEFAULT FALSE,
    client_status     VARCHAR(20) DEFAULT 'active' CHECK (client_status IN ('active', 'inactive', 'suspended')),
    notes             TEXT        DEFAULT ' ',
    created_at        TIMESTAMP   DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE installment
(
    id                UUID PRIMARY KEY,
    client_id         UUID           NOT NULL,
    total_amount      DECIMAL(10, 2) NOT NULL CHECK (total_amount >= 0),
    amount_paid       DECIMAL(10, 2) DEFAULT 0 CHECK (amount_paid >= 0),
    last_payment_date DATE,
    is_fully_paid     BOOLEAN GENERATED ALWAYS AS (total_amount <= amount_paid) STORED,
    currency_code     CHAR(3)        NOT NULL CHECK (currency_code IN ('usd', 'uzs')),
    created_at        TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    company_id        UUID           NOT NULL
);

-- Таблица payments
CREATE TABLE payments
(
    id             UUID PRIMARY KEY,
    installment_id UUID           NOT NULL REFERENCES installment (id) ON DELETE CASCADE,
    payment_date   TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    payment_amount DECIMAL(10, 2) NOT NULL CHECK (payment_amount > 0),
    created_at     TIMESTAMP   DEFAULT CURRENT_TIMESTAMP
);

-- Индексы для оптимизации запросов
CREATE INDEX idx_installment_client_id ON installment (client_id);