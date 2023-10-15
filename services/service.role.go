package services

import (
	"github.com/nutwreck/admin-pos-service/entities"
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type serviceRole struct {
	role entities.EntityRole
}

func NewServiceRole(role entities.EntityRole) *serviceRole {
	return &serviceRole{role: role}
}

/**
* ============================================
* Service Create New Master Role Teritory
*=============================================
 */

func (s *serviceRole) EntityCreate(input *schemes.Role) (*models.Role, schemes.SchemeDatabaseError) {
	var role schemes.Role
	role.Name = input.Name
	role.Type = input.Type
	role.MerchantID = input.MerchantID

	res, err := s.role.EntityCreate(&role)
	return res, err
}

/**
* =============================================
* Service Results All Master Role Teritory
*==============================================
 */

func (s *serviceRole) EntityResults(input *schemes.Role) (*[]schemes.GetAllRole, int64, schemes.SchemeDatabaseError) {
	var role schemes.Role
	role.Sort = input.Sort
	role.Page = input.Page
	role.PerPage = input.PerPage
	role.MerchantID = input.MerchantID
	role.Type = input.Type
	role.Name = input.Name
	role.ID = input.ID

	res, totalData, err := s.role.EntityResults(&role)
	return res, totalData, err
}

/**
* ==============================================
* Service Delete Master Role By ID Teritory
*===============================================
 */

func (s *serviceRole) EntityDelete(input *schemes.Role) (*models.Role, schemes.SchemeDatabaseError) {
	var role schemes.Role
	role.ID = input.ID

	res, err := s.role.EntityDelete(&role)
	return res, err
}

/**
* ==============================================
* Service Update Master Role By ID Teritory
*===============================================
 */

func (s *serviceRole) EntityUpdate(input *schemes.Role) (*models.Role, schemes.SchemeDatabaseError) {
	var role schemes.Role
	role.ID = input.ID
	role.Name = input.Name
	role.Type = input.Type
	role.MerchantID = input.MerchantID
	role.Active = input.Active

	res, err := s.role.EntityUpdate(&role)
	return res, err
}
