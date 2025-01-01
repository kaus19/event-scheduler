-- name: AddEventParticipant :exec
INSERT INTO event_participants (event_id, user_id, can_attend)
VALUES ($1, $2, $3)
ON CONFLICT (event_id, user_id) DO NOTHING;

-- name: GetEventParticipants :many
SELECT event_id, user_id, can_attend
FROM event_participants
WHERE event_id = $1;

-- name: GetParticipantStatus :one
SELECT event_id, user_id, can_attend
FROM event_participants
WHERE event_id = $1 AND user_id = $2;

-- name: UpdateAttendanceStatus :exec
UPDATE event_participants
SET can_attend = $3
WHERE event_id = $1 AND user_id = $2;

-- name: RemoveEventParticipant :exec
DELETE FROM event_participants
WHERE event_id = $1 AND user_id = $2;

-- name: ListUserEvents :many
SELECT event_id, user_id, can_attend
FROM event_participants
WHERE user_id = $1
ORDER BY event_id;
