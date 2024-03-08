-- +goose Up
SELECT
    'up SQL query';

CREATE TABLE
    holidays_jp (
        holiday_date TIMESTAMP NOT NULL,
        holiday_name TEXT NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    );

-- +goose Down
SELECT
    'down SQL query';

DROP TABLE holidays_jp;