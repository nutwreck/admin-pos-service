package entities

import (
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type EntitySupplier interface {
	EntityCreate(input *schemes.Supplier) (*models.Supplier, schemes.SchemeDatabaseError)
	EntityResults(input *schemes.Supplier) (*[]schemes.GetSupplier, int64, schemes.SchemeDatabaseError)
	EntityResult(input *schemes.Supplier) (*models.Supplier, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.Supplier) (*models.Supplier, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.Supplier) (*models.Supplier, schemes.SchemeDatabaseError)
}
