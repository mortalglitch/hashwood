-- +goose Up
CREATE TABLE files(
    id UUID PRIMARY KEY,
    file_name TEXT NOT NULL,
    directory TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    last_change TIMESTAMP NOT NULL,
    hash TEXT NOT NULL
);

CREATE TABLE history(
    id UUID PRIMARY KEY,
    previous_hash TEXT NOT NULL,
    current_hash TEXT NOT NULL,
    date_change TIMESTAMP NOT NULL,
    file_id UUID NOT NULL,
    CONSTRAINT fk_file_id
    FOREIGN KEY (file_id)
    REFERENCES files(id)
    ON DELETE CASCADE
);

CREATE TABLE ignorelist(
    id UUID PRIMARY KEY,
    file_name TEXT NOT NULL,
    directory TEXT NOT NULL,
    date_added TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE ignorelist;
DROP TABLE history;
DROP TABLE files;
