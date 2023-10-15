package entities

import (
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type EntityUnitOfMeasurement interface {
	EntityCreate(input *schemes.UnitOfMeasurement) (*models.UnitOfMeasurement, schemes.SchemeDatabaseError)
	EntityResults(input *schemes.UnitOfMeasurement) (*[]schemes.GetUnitOfMeasurement, int64, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.UnitOfMeasurement) (*models.UnitOfMeasurement, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.UnitOfMeasurement) (*models.UnitOfMeasurement, schemes.SchemeDatabaseError)
}
