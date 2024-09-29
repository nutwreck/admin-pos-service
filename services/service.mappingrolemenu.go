package services

import (
	"github.com/nutwreck/admin-pos-service/entities"
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type serviceMappingRoleMenu struct {
	mappingRoleMenu entities.EntityMappingRoleMenu
}

func NewServiceMappingRoleMenu(mappingRoleMenu entities.EntityMappingRoleMenu) *serviceMappingRoleMenu {
	return &serviceMappingRoleMenu{mappingRoleMenu: mappingRoleMenu}
}

/**
* ============================================
* Service Create New Master Mapping Role Menu Teritory
*=============================================
 */

func (s *serviceMappingRoleMenu) EntityCreate(input *schemes.MappingRoleMenu) (*models.MappingRoleMenu, schemes.SchemeDatabaseError) {
	var mappingRoleMenu schemes.MappingRoleMenu
	mappingRoleMenu.Name = input.Name
	mappingRoleMenu.MerchantID = input.MerchantID
	mappingRoleMenu.RoleID = input.RoleID
	mappingRoleMenu.MenuID = input.MenuID
	mappingRoleMenu.MenuDetailID = input.MenuDetailID
	mappingRoleMenu.MenuDetailFunctionID = input.MenuDetailFunctionID

	res, err := s.mappingRoleMenu.EntityCreate(&mappingRoleMenu)
	return res, err
}

/**
* =============================================
* Service Results All Master Mapping Role Menu Teritory
*==============================================
 */

func (s *serviceMappingRoleMenu) EntityResults(input *schemes.MappingRoleMenu) (*[]schemes.GetMappingRoleMenu, int64, schemes.SchemeDatabaseError) {
	var mappingRoleMenu schemes.MappingRoleMenu
	mappingRoleMenu.Sort = input.Sort
	mappingRoleMenu.Page = input.Page
	mappingRoleMenu.PerPage = input.PerPage
	mappingRoleMenu.MerchantID = input.MerchantID
	mappingRoleMenu.RoleID = input.RoleID
	mappingRoleMenu.MenuID = input.MenuID
	mappingRoleMenu.MenuDetailID = input.MenuDetailID
	mappingRoleMenu.MenuDetailFunctionID = input.MenuDetailFunctionID
	mappingRoleMenu.Name = input.Name
	mappingRoleMenu.ID = input.ID

	res, totalData, err := s.mappingRoleMenu.EntityResults(&mappingRoleMenu)
	return res, totalData, err
}

/**
* ==============================================
* Service Delete Master Mapping Role Menu By ID Teritory
*===============================================
 */

func (s *serviceMappingRoleMenu) EntityDelete(input *schemes.MappingRoleMenu) (*models.MappingRoleMenu, schemes.SchemeDatabaseError) {
	var mappingRoleMenu schemes.MappingRoleMenu
	mappingRoleMenu.ID = input.ID

	res, err := s.mappingRoleMenu.EntityDelete(&mappingRoleMenu)
	return res, err
}

/**
* ==============================================
* Service Update Master Mapping Role Menu By ID Teritory
*===============================================
 */

func (s *serviceMappingRoleMenu) EntityUpdate(input *schemes.MappingRoleMenu) (*models.MappingRoleMenu, schemes.SchemeDatabaseError) {
	var mappingRoleMenu schemes.MappingRoleMenu
	mappingRoleMenu.ID = input.ID
	mappingRoleMenu.MerchantID = input.MerchantID
	mappingRoleMenu.RoleID = input.RoleID
	mappingRoleMenu.MenuID = input.MenuID
	mappingRoleMenu.MenuDetailID = input.MenuDetailID
	mappingRoleMenu.MenuDetailFunctionID = input.MenuDetailFunctionID
	mappingRoleMenu.Name = input.Name
	mappingRoleMenu.Active = input.Active

	res, err := s.mappingRoleMenu.EntityUpdate(&mappingRoleMenu)
	return res, err
}
