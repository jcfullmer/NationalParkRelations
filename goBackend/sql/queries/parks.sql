-- name: GetParks :many
SELECT full_name, designation
FROM parks
ORDER BY full_name ASC;

-- name: GetParkByState :many
SELECT p.full_name, park_states.state_code
FROM parks as p
JOIN park_states
ON p.id = park_states.park_id
WHERE park_states.state_code = $1;
