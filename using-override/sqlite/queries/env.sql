-- name: EnvCreate :one
INSERT INTO env (
    id,
    create_time
) VALUES (
    ?,
    ?
)
RETURNING id, create_time;

-- name: EnvGet :one
SELECT * FROM env
WHERE id = ?;

-- See https://docs.sqlc.dev/en/latest/howto/named_parameters.html#nullable-parameters
-- The goal is to type create_time as *CreateTime for updates

-- This one doesn't have a nullable CreateTime
-- name: EnvUpdate :exec
UPDATE env SET
    create_time = COALESCE(sqlc.narg('create_time'), create_time)
WHERE id = sqlc.arg('id');

-- This one types it as *string
-- name: EnvUpdate2 :exec
UPDATE env SET
    create_time = COALESCE(CAST(sqlc.narg('create_time') AS TEXT), create_time)
WHERE id = sqlc.arg('id');

-- name: EnvUpdate2 :exec
UPDATE env SET
    create_time = COALESCE(CAST(sqlc.narg('create_time') AS TEXT), create_time)
WHERE id = sqlc.arg('id');