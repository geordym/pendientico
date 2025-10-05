-- +goose Up
-- +goose StatementBegin
CREATE TABLE workspace_contacts(
    id UUID PRIMARY KEY,
    workspace_id UUID NOT NULL REFERENCES workspaces(id) ON DELETE CASCADE,
    name VARCHAR(500) UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE workspace_contacts;
-- +goose StatementEnd
