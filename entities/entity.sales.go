package entities

import (
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type EntitySales interface {
	EntityCreate(input *schemes.Sales) (*models.Sales, schemes.SchemeDatabaseError)
	EntityResults(input *schemes.Sales) (*[]schemes.GetSales, int64, schemes.SchemeDatabaseError)
	EntityResult(input *schemes.Sales) (*models.Sales, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.Sales) (*models.Sales, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.Sales) (*models.Sales, schemes.SchemeDatabaseError)
}
