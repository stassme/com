CREATE TABLE legal_entities (
    id SERIAL PRIMARY KEY,
    bik VARCHAR(255),
    bank VARCHAR(255),
    address VARCHAR(255),
    kBill VARCHAR(255),
    rBill VARCHAR(255),
    currency VARCHAR(255),
    comment VARCHAR(255)
);