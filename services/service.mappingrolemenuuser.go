package services

import (
	"github.com/nutwreck/admin-pos-service/entities"
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type serviceMappingRoleMenuUser struct {
	mappingRoleMenuUser entities.EntityMappingRoleMenuUser
}

func NewServiceMappingRoleMenuUser(mappingRoleMenuUser entities.EntityMappingRoleMenuUser) *serviceMappingRoleMenuUser {
	return &serviceMappingRoleMenuUser{mappingRoleMenuUser: mappingRoleMenuUser}
}

/**
* ============================================
* Service Create New Mapping Role Menu User Teritory
*=============================================
 */

func (s *serviceMappingRoleMenuUser) EntityCreate(inputs *[]schemes.MappingRoleMenuUser) (*models.MappingRoleMenuUser, schemes.SchemeDatabaseError) {
	var createdMappingRoleMenuUser []schemes.MappingRoleMenuUser

	// Loop through each input in the batch
	for _, input := range *inputs {
		var role schemes.MappingRoleMenuUser
		role.MerchantID = input.MerchantID
		role.UserID = input.UserID
		role.MappingRoleMenuID = input.MappingRoleMenuID

		createdMappingRoleMenuUser = append(createdMappingRoleMenuUser, role)
	}

	res, err := s.mappingRoleMenuUser.EntityCreate(&createdMappingRoleMenuUser)
	return res, err
}

/**
* =============================================
* Service Results All Mapping Role Menu User Teritory
*==============================================
 */

func (s *serviceMappingRoleMenuUser) EntityResults(input *schemes.MappingRoleMenuUser) (*[]schemes.GetMappingRoleMenuUser, int64, schemes.SchemeDatabaseError) {
	var mappingRoleMenuUser schemes.MappingRoleMenuUser
	mappingRoleMenuUser.Sort = input.Sort
	mappingRoleMenuUser.Page = input.Page
	mappingRoleMenuUser.PerPage = input.PerPage
	mappingRoleMenuUser.MerchantID = input.MerchantID
	mappingRoleMenuUser.UserID = input.UserID
	mappingRoleMenuUser.MappingRoleMenuID = input.MappingRoleMenuID
	mappingRoleMenuUser.ID = input.ID

	res, totalData, err := s.mappingRoleMenuUser.EntityResults(&mappingRoleMenuUser)
	return res, totalData, err
}

/**
* ==============================================
* Service Delete Mapping Role Menu User By ID Teritory
*===============================================
 */

func (s *serviceMappingRoleMenuUser) EntityDelete(input *schemes.MappingRoleMenuUser) (*models.MappingRoleMenuUser, schemes.SchemeDatabaseError) {
	var mappingRoleMenuUser schemes.MappingRoleMenuUser
	mappingRoleMenuUser.ID = input.ID

	res, err := s.mappingRoleMenuUser.EntityDelete(&mappingRoleMenuUser)
	return res, err
}

/**
* ==============================================
* Service Update Mapping Role Menu User By ID Teritory
*===============================================
 */

func (s *serviceMappingRoleMenuUser) EntityUpdate(input *schemes.MappingRoleMenuUser) (*models.MappingRoleMenuUser, schemes.SchemeDatabaseError) {
	var mappingRoleMenuUser schemes.MappingRoleMenuUser
	mappingRoleMenuUser.ID = input.ID
	mappingRoleMenuUser.MappingRoleMenuID = input.MappingRoleMenuID
	mappingRoleMenuUser.UserID = input.UserID
	mappingRoleMenuUser.MerchantID = input.MerchantID
	mappingRoleMenuUser.Active = input.Active

	res, err := s.mappingRoleMenuUser.EntityUpdate(&mappingRoleMenuUser)
	return res, err
}
