-- name: CreateTimePreference :one
INSERT INTO time_preferences (owner_type, owner_id, start_time, end_time)
VALUES ($1, $2, $3, $4)
RETURNING id, owner_type, owner_id, start_time, end_time;

-- name: GetTimePreferenceByID :one
SELECT id, owner_type, owner_id, start_time, end_time
FROM time_preferences
WHERE id = $1;

-- name: GetTimePreferencesByOwner :many
SELECT id, owner_type, owner_id, start_time, end_time
FROM time_preferences
WHERE owner_type = $1 AND owner_id = $2
ORDER BY start_time;

-- name: UpdateTimePreference :exec
UPDATE time_preferences
SET start_time = $3, end_time = $4
WHERE owner_type = $1 AND owner_id = $2;

-- name: DeleteTimePreferenceByID :exec
DELETE FROM time_preferences
WHERE id = $1;

-- name: DeleteTimePreferenceByUnique :exec
DELETE FROM time_preferences
WHERE owner_type = $1 AND owner_id = $2 AND start_time = $3 AND end_time = $4;
