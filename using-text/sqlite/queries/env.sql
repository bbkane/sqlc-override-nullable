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
-- name: EnvUpdate :exec
UPDATE env SET
    create_time = COALESCE(sqlc.narg('create_time'), create_time)
WHERE id = sqlc.arg('id');
