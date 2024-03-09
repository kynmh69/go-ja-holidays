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

-- updated at function
CREATE FUNCTION refresh_updated_at_step1() RETURNS trigger AS
$$
BEGIN
  IF NEW.updated_at = OLD.updated_at THEN
    NEW.updated_at := NULL;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
    
CREATE FUNCTION refresh_updated_at_step2() RETURNS trigger AS
$$
BEGIN
  IF NEW.updated_at IS NULL THEN
    NEW.updated_at := OLD.updated_at;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE FUNCTION refresh_updated_at_step3() RETURNS trigger AS
$$
BEGIN
  IF NEW.updated_at IS NULL THEN
    NEW.updated_at := clock_timestamp();
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- create table
CREATE TABLE holidays_jp (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v7(),
    holiday_date TIMESTAMPTZ NOT NULL,
    holiday_name TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp()
);

CREATE TRIGGER refresh_holidays_jp_updated_at_step1
  BEFORE UPDATE ON holidays_jp FOR EACH ROW
  EXECUTE PROCEDURE refresh_updated_at_step1();
CREATE TRIGGER refresh_holidays_jp_updated_at_step2
  BEFORE UPDATE OF updated_at ON holidays_jp FOR EACH ROW
  EXECUTE PROCEDURE refresh_updated_at_step2();
CREATE TRIGGER refresh_holidays_jp_updated_at_step3
  BEFORE UPDATE ON holidays_jp FOR EACH ROW
  EXECUTE PROCEDURE refresh_updated_at_step3();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS holidays_jp;
-- +goose StatementEnd
