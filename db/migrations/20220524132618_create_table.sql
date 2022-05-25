-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS car (
    id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    car_name VARCHAR(30) NOT NULL,
    car_brand VARCHAR(30) NOT NULL,
    car_year VARCHAR(50)
) ENGINE INNODB;
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS car;
-- +goose StatementEnd