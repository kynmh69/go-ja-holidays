-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

-- UUID v7 function
CREATE OR REPLACE FUNCTION uuid_generate_v7() RETURNS UUID AS
$$
BEGIN
    return encode(set_bit(set_bit(overlay(
        uuid_send(gen_random_uuid())
        placing substring(int8send(floor(extract(epoch from clock_timestamp()) * 1000)::bigint) from 3)
        from 1 for 6
    ), 52, 1), 53, 1), 'hex')::uuid;
END
$$ LANGUAGE plpgsql;

-- create table
CREATE TABLE api_key (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v7(),
    api_key UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
    UNIQUE (api_key)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE api_key;
-- +goose StatementEnd
