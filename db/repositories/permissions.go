package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type PermissionRepository interface {
	GetPermissionById(id int64) (*models.Permissions, error)
	GetPermissionByName(name string) (*models.Permissions, error)
	GetAllPermissions() ([]*models.Permissions, error)
	CreatePermission(name string, description string) (*models.Permissions, error)
	DeletePermission(id int64) error
	UpdatePermission(id int64, name string, description string) (*models.Permissions, error)
}

type PermissionRepositoryImpl struct {
	db *sql.DB
}

func NewPermissionRepository(_db *sql.DB) PermissionRepository {
	return &PermissionRepositoryImpl{
		db: _db,
	}
}

func (r *PermissionRepositoryImpl) GetPermissionById(id int64) (*models.Permissions, error) {
	query := "SELECT id, name, description, resource, action, created_at, updated_at FROM permissions WHERE id = ?"
	row := r.db.QueryRow(query, id)

	permission := models.Permissions{}
	err := row.Scan(&permission.Id, &permission.Name, &permission.Description, &permission.Resource, &permission.Action, &permission.CreatedAt, &permission.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("permission with id %d not found", id)
		}
		return nil, err
	}

	return &permission, nil
}		

func (r *PermissionRepositoryImpl) GetPermissionByName(name string) (*models.Permissions, error) {
	query := "SELECT id, name, description, resource, action, created_at, updated_at FROM permissions WHERE name = ?"
	row := r.db.QueryRow(query, name)

	permission := models.Permissions{}
	err := row.Scan(&permission.Id, &permission.Name, &permission.Description, &permission.Resource, &permission.Action, &permission.CreatedAt, &permission.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("permission with name %s not found", name)
		}
		return nil, err
	}

	return &permission, nil
}	

func (r *PermissionRepositoryImpl) GetAllPermissions() ([]*models.Permissions, error) {
	query := "SELECT id, name, description, resource, action, created_at, updated_at FROM permissions"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []*models.Permissions
	for rows.Next() {
		permission := models.Permissions{}
		err := rows.Scan(&permission.Id, &permission.Name, &permission.Description, &permission.Resource, &permission.Action, &permission.CreatedAt, &permission.UpdatedAt)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, &permission)
	}

	return permissions, nil
}	

func (r *PermissionRepositoryImpl) CreatePermission(name string, description string) (*models.Permissions, error) {			
	query := "INSERT INTO permissions (name, description) VALUES (?, ?)"
	result, err := r.db.Exec(query, name, description)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return r.GetPermissionById(id)
}

func (r *PermissionRepositoryImpl) DeletePermission(id int64) error {
	query := "DELETE FROM permissions WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}			

func (r *PermissionRepositoryImpl) UpdatePermission(id int64, name string, description string) (*models.Permissions, error) {
	query := "UPDATE permissions SET name = ?, description = ?, updated_at = datetime('now') WHERE id = ?"
	_, err := r.db.Exec(query, name, description, id)
	if err != nil {
		return nil, err
	}

	return r.GetPermissionById(id)
}

