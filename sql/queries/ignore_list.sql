CREATE TABLE ignorelist(
    id UUID PRIMARY KEY,
    file_name TEXT NOT NULL,
    directory TEXT NOT NULL,
    date_added TIMESTAMP NOT NULL
);

-- name: CreateIgnoreEntry :one
INSERT INTO ignorelist (id, file_name, directory, date_added)
VALUES (
  $1,
  $2,
  $3,
  $4
)
RETURNING *;

-- name: DeleteIgnoreList :exec
DELETE FROM ignorelist;

-- name: GetIgnoreListByDateAdded :many
SELECT * FROM ignorelist
ORDER BY date_added;

-- name: DeleteIgnoreItemByID :exec
DELETE FROM ignorelist
WHERE ID = $1;

-- name: GetIgnoredItemByNameDirectory :one
SELECT * FROM ignorelist
WHERE file_name = $1 AND directory = $2;
