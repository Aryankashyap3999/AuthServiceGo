package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
	"strings"
)

type UserRoleRepository interface {
	// Define methods for user-role relationships here
	GetUserRoles(userId int64) ([]*models.Role, error)
	AssignRoleToUser(userId int64, roleId int64) error
	RemoveRoleFromUser(userId int64, roleId int64) error
	GetUserPermissions(userId int64) ([]*models.Permissions, error)
	HasPermission(userId int64, resource string, action string) (bool, error)
	HasRole(userId int64, roleName string) (bool, error)
	HasAllRoles(userId int64, roleNames []string) (bool, error)
	HasAnyRole(userId int64, roleNames []string) (bool, error)
}

type UserRoleRepositoryImpl struct {
	db *sql.DB
}

func NewUserRoleRepository(_db *sql.DB) UserRoleRepository {
	return &UserRoleRepositoryImpl{
		db: _db,
	}
}

// GetUserRoles retrieves all roles assigned to a user
func (urr *UserRoleRepositoryImpl) GetUserRoles(userId int64) ([]*models.Role, error) {
	query := `
		SELECT r.id, r.name, r.description, r.created_at, r.updated_at
		FROM roles r
		INNER JOIN user_roles ur ON r.id = ur.role_id
		WHERE ur.user_id = ?
	`

	rows, err := urr.db.Query(query, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to query user roles: %w", err)
	}
	defer rows.Close()

	var roles []*models.Role
	for rows.Next() {
		role := &models.Role{}
		err := rows.Scan(&role.Id, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan role: %w", err)
		}
		roles = append(roles, role)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating roles: %w", err)
	}

	return roles, nil
}

// AssignRoleToUser assigns a role to a user
func (urr *UserRoleRepositoryImpl) AssignRoleToUser(userId int64, roleId int64) error {
	query := `
		INSERT INTO user_roles (user_id, role_id, created_at, updated_at)
		VALUES (?, ?, NOW(), NOW())
		ON DUPLICATE KEY UPDATE updated_at = NOW()
	`

	_, err := urr.db.Exec(query, userId, roleId)
	if err != nil {
		return fmt.Errorf("failed to assign role to user: %w", err)
	}

	return nil
}

// RemoveRoleFromUser removes a role from a user
func (urr *UserRoleRepositoryImpl) RemoveRoleFromUser(userId int64, roleId int64) error {
	query := `
		DELETE FROM user_roles
		WHERE user_id = ? AND role_id = ?
	`

	result, err := urr.db.Exec(query, userId, roleId)
	if err != nil {
		return fmt.Errorf("failed to remove role from user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no role-user mapping found to delete")
	}

	return nil
}

// GetUserPermissions retrieves all permissions for a user through their roles
func (urr *UserRoleRepositoryImpl) GetUserPermissions(userId int64) ([]*models.Permissions, error) {
	query := `
		SELECT DISTINCT p.id, p.name, p.description, p.resource, p.action, p.created_at, p.updated_at
		FROM permissions p
		INNER JOIN role_permissions rp ON p.id = rp.permission_id
		INNER JOIN roles r ON rp.role_id = r.id
		INNER JOIN user_roles ur ON r.id = ur.role_id
		WHERE ur.user_id = ?
	`

	rows, err := urr.db.Query(query, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to query user permissions: %w", err)
	}
	defer rows.Close()

	var permissions []*models.Permissions
	for rows.Next() {
		permission := &models.Permissions{}
		err := rows.Scan(&permission.Id, &permission.Name, &permission.Description,
			&permission.Resource, &permission.Action, &permission.CreatedAt, &permission.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan permission: %w", err)
		}
		permissions = append(permissions, permission)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating permissions: %w", err)
	}

	return permissions, nil
}

// HasPermission checks if a user has a specific permission on a resource
func (urr *UserRoleRepositoryImpl) HasPermission(userId int64, resource string, action string) (bool, error) {
	query := `
		SELECT COUNT(*) as count
		FROM permissions p
		INNER JOIN role_permissions rp ON p.id = rp.permission_id
		INNER JOIN roles r ON rp.role_id = r.id
		INNER JOIN user_roles ur ON r.id = ur.role_id
		WHERE ur.user_id = ? AND p.resource = ? AND p.action = ?
	`

	var count int
	err := urr.db.QueryRow(query, userId, resource, action).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to check permission: %w", err)
	}

	return count > 0, nil
}

// HasRole checks if a user has a specific role
func (urr *UserRoleRepositoryImpl) HasRole(userId int64, roleName string) (bool, error) {
	query := `
		SELECT COUNT(*) as count
		FROM roles r
		INNER JOIN user_roles ur ON r.id = ur.role_id
		WHERE ur.user_id = ? AND r.name = ?
	`

	var count int
	err := urr.db.QueryRow(query, userId, roleName).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to check role: %w", err)
	}

	return count > 0, nil
}

func (urr *UserRoleRepositoryImpl) HasAllRoles(userId int64, roleNames []string) (bool, error) {

	if len(roleNames) == 0 {	
		return true, nil
	}

	query := `
		SELECT COUNT(*) = ?
		FROM user_roles ur
		INNER JOIN roles r ON ur.role_id = r.id
		WHERE ur.user_id = ? AND r.name IN (?)
		GROUP BY ur.user_id
		`

	row := urr.db.QueryRow(query, len(roleNames), userId, strings.Join(roleNames, ","))

	var hasAll bool
	err := row.Scan(&hasAll)	
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("failed to check all roles: %w", err)
	}

	return hasAll, nil
}	

func (urr *UserRoleRepositoryImpl) HasAnyRole(userId int64, roleNames []string) (bool, error) {

	if len(roleNames) == 0 {	
		return false, nil
	}

	placeholders := strings.Repeat("?,", len(roleNames))
	placeholders = placeholders[:len(placeholders)-1] // Remove trailing comma

	query := fmt.Sprintf(` 
		SELECT COUNT(*) > 0
		FROM user_roles ur
		INNER JOIN roles r ON ur.role_id = r.id
		WHERE ur.user_id = ? AND r.name IN (%s)
		`, placeholders)

	// roleNameStr := utils.FormateRoles(roleNames)	

	args := make([]interface{}, 0, 1 + len(roleNames))
	args = append(args, userId)
	for _, roleName := range roleNames {
		args = append(args, roleName)
	}	

	row := urr.db.QueryRow(query, args...)

	var hasAny bool
	err := row.Scan(&hasAny)	
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("failed to check any roles: %w", err)
	}

	return hasAny, nil
}	


