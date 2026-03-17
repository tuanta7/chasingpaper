-- name: ListPlans :many
SELECT * FROM plan ORDER BY id;

-- name: GetPlan :one
SELECT * FROM plan WHERE id = $1 LIMIT 1;

-- name: CreatePlan :exec
INSERT INTO plan (name, description)
VALUES ($1, $2);

-- name: UpdatePlan :exec
UPDATE plan
SET
    name = $1,
    description = $2
WHERE id = $3
RETURNING *;

-- name: DeletePlan :exec
DELETE FROM plan WHERE id = $1;