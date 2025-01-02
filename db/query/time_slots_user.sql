-- name: CreateTimeSlotUser :one
INSERT INTO time_slots_user (user_id, start_time, end_time)
VALUES ($1, $2, $3)
RETURNING id, user_id, start_time, end_time;

-- name: GetTimePreferencesByUser :many
SELECT id, user_id, start_time, end_time
FROM time_slots_user
WHERE user_id = $1
ORDER BY start_time;

-- name: GetTimePreferencesForAllUsers :many
SELECT user_id, start_time, end_time
FROM time_slots_user;

-- name: UpdateTimePreferenceUser :exec
UPDATE time_slots_user
SET start_time = $3, end_time = $4
WHERE id = $1 AND user_id = $2;

-- name: DeleteTimePreferenceUser :exec
DELETE FROM time_slots_user
WHERE id = $1 AND user_id = $2;
