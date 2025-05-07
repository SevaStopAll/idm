CREATE TABLE employee
(
    id         bigint,
    name       text,
    created_at timestamptz,
    updated_at timestamptz
)

CREATE TABLE role
(
    id         bigint,
    name       text,
    created_at timestamptz,
    updated_at timestamptz
)

CREATE TABLE employee_role
(
    id          bigint,
    employee_id bigint references employee.id,
    role_id     bigint references role.id,
    created_at timestamptz,
)