package entities

import (
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type EntityProduct interface {
	EntityCreate(input *[]schemes.Product) (*models.Product, schemes.SchemeDatabaseError)
	EntityResults(input *schemes.Product) (*[]schemes.GetProduct, int64, schemes.SchemeDatabaseError)
	EntityResult(input *schemes.Product) (*models.Product, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.Product) (*models.Product, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.Product) (*models.Product, schemes.SchemeDatabaseError)
}
