-- +goose Up
-- +goose StatementBegin
CREATE TABLE employee
(
    id         bigint primary key GENERATED ALWAYS AS IDENTITY,
    name       text not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now()
);

CREATE TABLE role
(
    id         bigint primary key GENERATED ALWAYS AS IDENTITY,
    name       text not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now()
);

CREATE TABLE employee_role
(
    id          bigint primary key GENERATED ALWAYS AS IDENTITY,
    employee_id bigint not null references employee(id),
    role_id     bigint not null references role(id),
    created_at timestamptz default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE employee_role;

DROP TABLE role;

DROP TABLE employee;
-- +goose StatementEnd
