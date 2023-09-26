package entities

import (
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type EntityMenuDetail interface {
	EntityCreate(input *schemes.MenuDetail) (*models.MenuDetail, schemes.SchemeDatabaseError)
	EntityResults(input *schemes.MenuDetail) (*[]schemes.GetMenuDetail, int64, schemes.SchemeDatabaseError)
	EntityResult(input *schemes.MenuDetail) (*models.MenuDetail, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.MenuDetail) (*models.MenuDetail, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.MenuDetail) (*models.MenuDetail, schemes.SchemeDatabaseError)
}
