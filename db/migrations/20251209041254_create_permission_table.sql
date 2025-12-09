-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS permissions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    resource VARCHAR(100) NOT NULL,
    action VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)

-- +goose StatementEnd

-- +goose StatementBegin

INSERT INTO permissions (name, description, resource, action) VALUES 
('user:read', 'Permission to read user information', 'user', 'read'),
('user:manage', 'Permission to manage user accounts', 'user', 'manage'),
('user:delete', 'Permission to delete user accounts', 'user', 'delete'),
('role:read', 'Permission to read role information', 'role', 'read'),
('role:write', 'Permission to create or update roles', 'role', 'write'),
('role:delete', 'Permission to delete roles', 'role', 'delete'),
('role:manage', 'Permission to manage roles', 'role', 'manage'),
('permission:read', 'Permission to read permission information', 'permission', 'read'),
('permission:write', 'Permission to create or update permissions', 'permission', 'write'),
('permission:delete', 'Permission to delete permissions', 'permission', 'delete'),
('permission:manage', 'Permission to manage permissions', 'permission', 'manage');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS permissions;
-- +goose StatementEnd
