package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type RolePermissionRepository interface {
	GetRolePermissionById(id int64) (*models.RolePermissions, error)
	GetRolePermissionByRoleId(roleId int64) ([]*models.RolePermissions, error)
	GetRolePermissionByPermissionId(permissionId int64) ([]*models.RolePermissions, error)
	AddPermissionToRole(roleId int64, permissionId int64) error
	RemovePermissionFromRole(roleId int64, permissionId int64) error
	GetAllRolePermissionsFromRole() ([]*models.RolePermissions, error)
}

type RolePermissionRepositoryImpl struct {
	db *sql.DB
}

func NewRolePermissionRepository(_db *sql.DB) RolePermissionRepository {
	return &RolePermissionRepositoryImpl{
		db: _db,
	}
}

// GetRolePermissionById retrieves a specific role-permission mapping by ID
func (rpr *RolePermissionRepositoryImpl) GetRolePermissionById(id int64) (*models.RolePermissions, error) {
	query := `
		SELECT id, role_id, permission_id, created_at, updated_at
		FROM role_permissions
		WHERE id = ?
	`

	rolePermission := &models.RolePermissions{}
	err := rpr.db.QueryRow(query, id).Scan(
		&rolePermission.Id,
		&rolePermission.RoleId,
		&rolePermission.PermissionId,
		&rolePermission.CreatedAt,
		&rolePermission.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("role-permission mapping not found with id: %d", id)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query role-permission: %w", err)
	}

	return rolePermission, nil
}

// GetRolePermissionByRoleId retrieves all permissions assigned to a role
func (rpr *RolePermissionRepositoryImpl) GetRolePermissionByRoleId(roleId int64) ([]*models.RolePermissions, error) {
	query := `
		SELECT id, role_id, permission_id, created_at, updated_at
		FROM role_permissions
		WHERE role_id = ?
	`

	rows, err := rpr.db.Query(query, roleId)
	if err != nil {
		return nil, fmt.Errorf("failed to query role-permissions by role: %w", err)
	}
	defer rows.Close()

	var rolePermissions []*models.RolePermissions
	for rows.Next() {
		rp := &models.RolePermissions{}
		err := rows.Scan(&rp.Id, &rp.RoleId, &rp.PermissionId, &rp.CreatedAt, &rp.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan role-permission: %w", err)
		}
		rolePermissions = append(rolePermissions, rp)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating role-permissions: %w", err)
	}

	return rolePermissions, nil
}

// GetRolePermissionByPermissionId retrieves all roles that have a specific permission
func (rpr *RolePermissionRepositoryImpl) GetRolePermissionByPermissionId(permissionId int64) ([]*models.RolePermissions, error) {
	query := `
		SELECT id, role_id, permission_id, created_at, updated_at
		FROM role_permissions
		WHERE permission_id = ?
	`

	rows, err := rpr.db.Query(query, permissionId)
	if err != nil {
		return nil, fmt.Errorf("failed to query role-permissions by permission: %w", err)
	}
	defer rows.Close()

	var rolePermissions []*models.RolePermissions
	for rows.Next() {
		rp := &models.RolePermissions{}
		err := rows.Scan(&rp.Id, &rp.RoleId, &rp.PermissionId, &rp.CreatedAt, &rp.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan role-permission: %w", err)
		}
		rolePermissions = append(rolePermissions, rp)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating role-permissions: %w", err)
	}

	return rolePermissions, nil
}

// AddPermissionToRole assigns a permission to a role
func (rpr *RolePermissionRepositoryImpl) AddPermissionToRole(roleId int64, permissionId int64) error {
	query := `
		INSERT INTO role_permissions (role_id, permission_id, created_at, updated_at)
		VALUES (?, ?, NOW(), NOW())
		ON DUPLICATE KEY UPDATE updated_at = NOW()
	`

	_, err := rpr.db.Exec(query, roleId, permissionId)
	if err != nil {
		return fmt.Errorf("failed to add permission to role: %w", err)
	}

	return nil
}

// RemovePermissionFromRole removes a permission from a role
func (rpr *RolePermissionRepositoryImpl) RemovePermissionFromRole(roleId int64, permissionId int64) error {
	query := `
		DELETE FROM role_permissions
		WHERE role_id = ? AND permission_id = ?
	`

	result, err := rpr.db.Exec(query, roleId, permissionId)
	if err != nil {
		return fmt.Errorf("failed to remove permission from role: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no role-permission mapping found to delete")
	}

	return nil
}

// GetAllRolePermissionsFromRole retrieves all role-permission mappings (no filter)
func (rpr *RolePermissionRepositoryImpl) GetAllRolePermissionsFromRole() ([]*models.RolePermissions, error) {
	query := `
		SELECT id, role_id, permission_id, created_at, updated_at
		FROM role_permissions
	`

	rows, err := rpr.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query all role-permissions: %w", err)
	}
	defer rows.Close()

	var rolePermissions []*models.RolePermissions
	for rows.Next() {
		rp := &models.RolePermissions{}
		err := rows.Scan(&rp.Id, &rp.RoleId, &rp.PermissionId, &rp.CreatedAt, &rp.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan role-permission: %w", err)
		}
		rolePermissions = append(rolePermissions, rp)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating role-permissions: %w", err)
	}

	return rolePermissions, nil
}
