-- +goose Up
-- +goose StatementBegin
CREATE TABLE employee
(
    id         bigint GENERATED ALWAYS AS IDENTITY,
    name       text,
    created_at timestamptz,
    updated_at timestamptz
);

CREATE TABLE role
(
    id         bigint GENERATED ALWAYS AS IDENTITY,
    name       text,
    created_at timestamptz,
    updated_at timestamptz
);

CREATE TABLE employee_role
(
    id          bigint GENERATED ALWAYS AS IDENTITY,
    employee_id bigint,
    role_id     bigint,
    created_at timestamptz
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE employee_role;

DROP TABLE role;

DROP TABLE employee;
-- +goose StatementEnd
