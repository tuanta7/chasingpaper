-- name: ListPlans :many
SELECT * FROM plan
ORDER BY created_at DESC
OFFSET $1 LIMIT $2;

-- name: GetPlan :one
SELECT * FROM plan WHERE id = $1 LIMIT 1;

-- name: CreatePlan :one
INSERT INTO plan (id, name, description, is_active)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdatePlan :one
UPDATE plan
SET
    name = $1,
    description = $2,
    is_active = $3
WHERE id = $4
RETURNING *;

-- name: DeletePlan :exec
DELETE FROM plan WHERE id = $1;

-- name: CreatePrice :one
INSERT INTO price (
    plan_id,
    provider,
    provider_plan_id,
    cached_response,
    last_synced_at
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetPricesByPlan :many
SELECT * FROM price WHERE plan_id = $1 ORDER BY created_at DESC;

-- name: DeletePrice :exec
DELETE FROM price WHERE plan_id = $1 AND provider = $2 AND provider_plan_id = $3;
