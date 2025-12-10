package services

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/models"
)

type RoleService interface {
	GetRoleById(id int64) (*models.Role, error)
	GetRoleByName(name string) (*models.Role, error)
	GetAllRoles() ([]*models.Role, error)
	CreateRoles(name string, description string) (*models.Role, error)
	DeleteRoleById(id int64) error
	UpdateRole(id int64, name string, description string) (*models.Role, error) 
	GetRolePermissions(roleId int64) ([]*models.RolePermissions, error)
	AddPermissionToRole(roleId int64, permissionId int64) error
}

type RoleServiceImpl struct {
	roleRepository db.RoleRepository
	rolePermissionRepository db.RolePermissionRepository
}

func NewRoleService(roleRepo db.RoleRepository) RoleService {
	return &RoleServiceImpl{
		roleRepository: roleRepo,
	}
}

func (r *RoleServiceImpl) GetRoleById(id int64) (*models.Role, error) {
	return r.roleRepository.GetRoleById(id)
}	

func (r *RoleServiceImpl) GetRoleByName(name string) (*models.Role, error) {
	return r.roleRepository.GetRoleByName(name)
}

func (r *RoleServiceImpl) GetAllRoles() ([]*models.Role, error) {
	return r.roleRepository.GetAllRoles()
}

func (r *RoleServiceImpl) CreateRoles(name string, description string) (*models.Role, error) {
	return r.roleRepository.CreateRoles(name, description)
}

func (r *RoleServiceImpl) DeleteRoleById(id int64) error {
	return r.roleRepository.DeleteRole(id)
}	

func (r *RoleServiceImpl) UpdateRole(id int64, name string, description string) (*models.Role, error) {
	return r.roleRepository.UpdateRole(id, name, description)
}

func (r *RoleServiceImpl) GetRolePermissions(roleId int64) ([]*models.RolePermissions, error) {
	return r.rolePermissionRepository.GetRolePermissionByRoleId(roleId)
}

func (r *RoleServiceImpl) AddPermissionToRole(roleId int64, permissionId int64) error {
	return r.rolePermissionRepository.AddPermissionToRole(roleId, permissionId)
}