-- +goose Up
CREATE TABLE IF NOT EXISTS plan (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS price (
    plan_id UUID NOT NULL REFERENCES plan(id) ON DELETE CASCADE,
    provider VARCHAR(50) NOT NULL,
    provider_plan_id VARCHAR(255) NOT NULL,
    cached_response JSONB,
    last_synced_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (plan_id, provider, provider_plan_id)
);

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS price;
DROP TABLE IF EXISTS plan;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
