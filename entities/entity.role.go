package entities

import (
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type EntityRole interface {
	EntityCreate(input *schemes.Role) (*models.Role, schemes.SchemeDatabaseError)
	EntityResults(input *schemes.Role) (*[]models.Role, int64, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.Role) (*models.Role, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.Role) (*models.Role, schemes.SchemeDatabaseError)
}
