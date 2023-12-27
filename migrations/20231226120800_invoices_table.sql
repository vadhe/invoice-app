-- +goose Up
-- +goose StatementBegin
CREATE TABLE invoices (
    id SERIAL,
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    id_invoices  VARCHAR(10) NOT NULL PRIMARY KEY,
    status VARCHAR(20) CHECK (status IN ('paid', 'draft', 'pending')) NOT NULL,
    price DECIMAL(10, 2), -- Adjust the precision and scale as needed
    name VARCHAR(50)
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS invoices;
-- +goose StatementEnd
