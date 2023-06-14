
-- +migrate Up
CREATE TABLE socks (
    id SERIAL PRIMARY KEY,
    color VARCHAR(255) NOT NULL,
    cotton_part INT NOT NULL,
    quantity INT NOT NULL
);
-- +migrate Down
DROP TABLE socks;