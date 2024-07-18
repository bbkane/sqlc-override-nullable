-- name: EnvCreate :one
INSERT INTO env (
    create_time
) VALUES (
    ?
)
RETURNING id, create_time;

-- name: EnvList :many
SELECT * FROM env
ORDER BY create_time ASC;

-- name: EnvList :one
SELECT * FROM env
WHERE id = ?;

-- See https://docs.sqlc.dev/en/latest/howto/named_parameters.html#nullable-parameters
-- name: EnvUpdate :exec
UPDATE env SET
    create_time = COALESCE(sqlc.narg('create_time'), create_time)
WHERE id = sqlc.arg('id');
