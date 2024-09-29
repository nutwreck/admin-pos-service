package entities

import (
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type EntityMappingRoleMenu interface {
	EntityCreate(input *schemes.MappingRoleMenu) (*models.MappingRoleMenu, schemes.SchemeDatabaseError)
	EntityResults(input *schemes.MappingRoleMenu) (*[]schemes.GetMappingRoleMenu, int64, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.MappingRoleMenu) (*models.MappingRoleMenu, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.MappingRoleMenu) (*models.MappingRoleMenu, schemes.SchemeDatabaseError)
}
