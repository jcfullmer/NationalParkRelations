-- +goose Up
CREATE TABLE states (
    code CHAR(2) PRIMARY KEY,
    name TEXT
);

CREATE TABLE park_states (
    park_id UUID REFERENCES parks(id) ON DELETE CASCADE,
    state_code CHAR(2) REFERENCES states(code) ON DELETE CASCADE,
    PRIMARY KEY (park_id, state_code)
);

-- +goose Down
DROP TABLE park_states;
DROP TABLE states;
