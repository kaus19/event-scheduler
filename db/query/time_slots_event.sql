-- name: CreateTimeSlotEvent :one
INSERT INTO time_slots_event (event_id, start_time, end_time)
VALUES ($1, $2, $3)
RETURNING id, event_id, start_time, end_time;

-- name: GetTimeSlotsByEvent :many
SELECT id, event_id, start_time, end_time
FROM time_slots_event
WHERE event_id = $1
ORDER BY start_time;

-- name: UpdateTimeSlotEvent :exec
UPDATE time_slots_event
SET start_time = $3, end_time = $4
WHERE id = $1 AND event_id = $2;

-- name: DeleteTimeSlotEvent :exec
DELETE FROM time_slots_event
WHERE id = $1 AND event_id = $2;