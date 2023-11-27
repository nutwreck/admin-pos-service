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

	EntityGetMenu(input *schemes.Menu) (*[]schemes.GetMenu, schemes.SchemeDatabaseError)
	EntityGetMenuDetail(input *schemes.MenuDetail) (*[]schemes.GetMenuDetail, schemes.SchemeDatabaseError)
	EntityGetMenuDetailFunction(input *schemes.MenuDetailFunction) (*[]schemes.GetMenuDetailFunction, schemes.SchemeDatabaseError)
}
