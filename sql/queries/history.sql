-- name: CreateHistoryEntry :one
INSERT INTO history (id, previous_hash, current_hash, date_change, file_id)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
)
RETURNING *;

-- name: DeleteHistory :exec
DELETE FROM history;

-- name: GetHistoryByDateChanged :many
SELECT * FROM history
ORDER BY date_change;

-- name: GetHistoryByFileID :many
SELECT * FROM history
WHERE file_id = $1;
