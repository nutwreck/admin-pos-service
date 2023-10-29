package entities

import (
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type EntityProductCategorySub interface {
	EntityCreate(input *schemes.ProductCategorySub) (*models.ProductCategorySub, schemes.SchemeDatabaseError)
	EntityResults(input *schemes.ProductCategorySub) (*[]schemes.GetProductCategorySub, int64, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.ProductCategorySub) (*models.ProductCategorySub, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.ProductCategorySub) (*models.ProductCategorySub, schemes.SchemeDatabaseError)
}
