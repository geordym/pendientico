-- +goose Up
-- +goose StatementBegin
CREATE TABLE workspaces (
    id UUID PRIMARY KEY,
    name VARCHAR(150)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE workspaces;
-- +goose StatementEnd
