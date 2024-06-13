-- Tabela bills
CREATE TABLE IF NOT EXISTS bills (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    amount VARCHAR(10) NOT NULL,
    installment BIGINT NOT NULL, 
);