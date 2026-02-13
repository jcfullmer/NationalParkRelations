-- +goose Up
INSERT INTO states(code)
SELECT DISTINCT trim(s)
FROM parks, unnest(string_to_array(states, ',')) AS s
ON CONFLICT DO NOTHING;

INSERT INTO park_states (park_id, state_code)
SELECT id, trim(s)
FROM parks, unnest(string_to_array(states, ',')) AS s;

ALTER TABLE parks
DROP COLUMN states;
