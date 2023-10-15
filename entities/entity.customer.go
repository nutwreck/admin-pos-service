package entities

import (
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type EntityCustomer interface {
	EntityCreate(input *schemes.Customer) (*models.Customer, schemes.SchemeDatabaseError)
	EntityResults(input *schemes.Customer) (*[]schemes.GetCustomer, int64, schemes.SchemeDatabaseError)
	EntityResult(input *schemes.Customer) (*models.Customer, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.Customer) (*models.Customer, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.Customer) (*models.Customer, schemes.SchemeDatabaseError)
}
