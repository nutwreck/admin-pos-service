package entities

import (
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type EntityMenu interface {
	EntityCreate(input *schemes.Menu) (*models.Menu, schemes.SchemeDatabaseError)
	EntityResults(input *schemes.Menu) (*[]models.Menu, int64, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.Menu) (*models.Menu, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.Menu) (*models.Menu, schemes.SchemeDatabaseError)
}
