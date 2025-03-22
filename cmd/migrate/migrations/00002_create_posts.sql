-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS posts (
    id bigserial PRIMARY KEY,
    content varchar(255) NOT NULL,
    title varchar(255) NOT NULL,
	user_id bigint NOT NULL,
    tags varchar(255) NOT NULL,
    create_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

-- +goose StatementEnd



-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE IF EXISTS posts;
-- +goose StatementEnd
