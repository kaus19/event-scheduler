-- name: CreateEvent :one
INSERT INTO events (organizer_id, event_name, event_description, duration)
VALUES ($1, $2, $3, $4)
RETURNING event_id, organizer_id, event_name, event_description, duration, created_at, updated_at;

-- name: GetEventByID :one
SELECT event_id, organizer_id, event_name, event_description, duration, created_at, updated_at
FROM events
WHERE event_id = $1;

-- name: ListEvents :many
SELECT event_id, organizer_id, event_name, event_description, duration, created_at, updated_at
FROM events
ORDER BY created_at DESC;

-- name: ListEventsByOrganizer :many
SELECT event_id, organizer_id, event_name, event_description, duration, created_at, updated_at
FROM events
WHERE organizer_id = $1
ORDER BY created_at DESC;

-- name: UpdateEvent :exec
UPDATE events
SET event_name = $2, event_description = $3, duration = $4, updated_at = $5
WHERE event_id = $1;

-- name: DeleteEvent :exec
DELETE FROM events
WHERE event_id = $1;
