package entities

import (
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type EntityMenuDetailFunction interface {
	EntityCreate(input *schemes.MenuDetailFunction) (*models.MenuDetailFunction, schemes.SchemeDatabaseError)
	EntityResults(input *schemes.MenuDetailFunction) (*[]schemes.GetMenuDetailFunction, int64, schemes.SchemeDatabaseError)
	EntityResult(input *schemes.MenuDetailFunction) (*models.MenuDetailFunction, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.MenuDetailFunction) (*models.MenuDetailFunction, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.MenuDetailFunction) (*models.MenuDetailFunction, schemes.SchemeDatabaseError)
}
