package entities

import (
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type EntityUnitOfMeasurementType interface {
	EntityCreate(input *schemes.UnitOfMeasurementType) (*models.UnitOfMeasurementType, schemes.SchemeDatabaseError)
	EntityResults(input *schemes.UnitOfMeasurementType) (*[]schemes.GetUnitOfMeasurementType, int64, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.UnitOfMeasurementType) (*models.UnitOfMeasurementType, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.UnitOfMeasurementType) (*models.UnitOfMeasurementType, schemes.SchemeDatabaseError)
}
