package entities

import (
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type EntityMenu interface {
	EntityCreate(input *schemes.Menu) (*models.Menu, schemes.SchemeDatabaseError)
	EntityResults(input *schemes.Menu) (*[]schemes.GetMenu, int64, schemes.SchemeDatabaseError)
	EntityResultRelations(input *schemes.Menu) (*[]schemes.GetMenuRelation, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.Menu) (*models.Menu, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.Menu) (*models.Menu, schemes.SchemeDatabaseError)
}
