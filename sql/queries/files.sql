-- name: CreateFileHash :one
INSERT INTO files (id, file_name, directory, created_at, updated_at, last_change, hash)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7
)
RETURNING *;

-- name: DeleteFiles :exec
DELETE FROM files;

-- name: GetFileHashByName :one
SELECT * FROM files
WHERE hash = $1 LIMIT 1;

-- name: UpdateFileChecked :exec
UPDATE files
SET updated_at = $1
WHERE id = $2;

-- name: UpdateFileHash :exec
UPDATE files
SET updated_at = $1, last_change = $1, hash = $2
WHERE id = $3;
