-- Simple migrations for demo. For production, use a dedicated migration tool (e.g. golang-migrate).

CREATE TABLE IF NOT EXISTS wallets (
    id SERIAL PRIMARY KEY,
    address VARCHAR(128) NOT NULL UNIQUE,
    public_key TEXT NOT NULL,
    private_key_encrypted TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    tx_hash VARCHAR(128) NOT NULL UNIQUE,
    from_address VARCHAR(128) NOT NULL,
    to_address VARCHAR(128),
    data JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);
