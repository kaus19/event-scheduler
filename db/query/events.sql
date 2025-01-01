-- name: CreateEvent :one
INSERT INTO events (organizer_id, event_name, event_description)
VALUES ($1, $2, $3)
RETURNING event_id, organizer_id, event_name, event_description, created_at;

-- name: GetEventByID :one
SELECT event_id, organizer_id, event_name, event_description, created_at
FROM events
WHERE event_id = $1;

-- name: ListEvents :many
SELECT event_id, organizer_id, event_name, event_description, created_at
FROM events
ORDER BY created_at DESC;

-- name: ListEventsByOrganizer :many
SELECT event_id, organizer_id, event_name, event_description, created_at
FROM events
WHERE organizer_id = $1
ORDER BY created_at DESC;

-- name: UpdateEventDescription :exec
UPDATE events
SET event_description = $2
WHERE event_id = $1;

-- name: DeleteEvent :exec
DELETE FROM events
WHERE event_id = $1;
