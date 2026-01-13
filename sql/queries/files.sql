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

-- name: GetFileByName :one
SELECT * FROM files
WHERE file_name = $1 AND directory = $2 LIMIT 1;

-- name: GetHashByID :one
SELECT hash FROM files
WHERE id = $1 LIMIT 1;

-- name: UpdateFileChecked :exec
UPDATE files
SET updated_at = $1
WHERE id = $2;

-- name: UpdateFileHash :exec
UPDATE files
SET updated_at = $1, last_change = $1, hash = $2
WHERE id = $3;
