package entities

import (
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type EntityMappingRoleMenuUser interface {
	EntityCreate(input *[]schemes.MappingRoleMenuUser) (*models.MappingRoleMenuUser, schemes.SchemeDatabaseError)
	EntityResults(input *schemes.MappingRoleMenuUser) (*[]schemes.GetMappingRoleMenuUser, int64, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.MappingRoleMenuUser) (*models.MappingRoleMenuUser, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.MappingRoleMenuUser) (*models.MappingRoleMenuUser, schemes.SchemeDatabaseError)
}
