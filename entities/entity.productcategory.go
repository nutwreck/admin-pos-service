package entities

import (
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type EntityProductCategory interface {
	EntityCreate(input *schemes.ProductCategory) (*models.ProductCategory, schemes.SchemeDatabaseError)
	EntityResults(input *schemes.ProductCategory) (*[]schemes.GetProductCategory, int64, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.ProductCategory) (*models.ProductCategory, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.ProductCategory) (*models.ProductCategory, schemes.SchemeDatabaseError)
}
