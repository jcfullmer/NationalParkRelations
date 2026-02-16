-- name: GetParks :many
SELECT full_name, designation
FROM parks
ORDER BY 
  CASE WHEN $1 = 'full_name' THEN parks.full_name END ASC,
  CASE WHEN $1 = 'designation' THEN parks.designation END ASC,
    parks.park_code ASC
LIMIT $2;

-- name: GetParkByState :many
SELECT p.full_name, park_states.state_code
FROM parks as p
JOIN park_states
ON p.id = park_states.park_id
WHERE park_states.state_code = $1
ORDER BY 
  CASE WHEN $1 = 'full_name' THEN p.full_name END ASC,
  CASE WHEN $1 = 'designation' THEN p.designation END ASC,
    p.id ASC
LIMIT $2;
