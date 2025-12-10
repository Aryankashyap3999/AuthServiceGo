package controllers

import (
	"AuthInGo/services"
	"AuthInGo/utils"
	"fmt"
	"net/http"
	"strconv"
)

type RoleController struct {
	RoleService services.RoleService
}

func NewRoleController(roleService services.RoleService) *RoleController {
	return &RoleController{
		RoleService: roleService,
	}
}

func (rc *RoleController) GetRoleById(w http.ResponseWriter, r *http.Request) {
	roleId := r.PathValue("id")

	if roleId == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Role ID is required", fmt.Errorf("role ID is missing"))
		return
	}

	id, err := strconv.ParseInt(roleId, 10, 64)

	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid Role ID", err)
		return 
	}

	role, err := rc.RoleService.GetRoleById(id)

	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Error fetching role", err)
		return
	}	

	if role == nil {
		utils.WriteJsonErrorResponse(w, http.StatusNotFound, "Role not found", nil)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Role fetched successfully", role)
}

func (rc *RoleController) GetAllRoles() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		roles, err := rc.RoleService.GetAllRoles()

		if err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Error fetching roles", err)
			return
		}

		utils.WriteJsonSuccessResponse(w, http.StatusOK, "Roles fetched successfully", roles)
	}
}

func (rc *RoleController) GetRolePermissions(w http.ResponseWriter, r *http.Request) {
	roles, err := rc.RoleService.GetAllRoles()

	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Error fetching roles", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Roles fetched successfully", roles)
}