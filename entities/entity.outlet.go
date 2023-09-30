package entities

import (
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type EntityOutlet interface {
	EntityCreate(input *schemes.Outlet) (*models.Outlet, schemes.SchemeDatabaseError)
	EntityResults(input *schemes.Outlet) (*[]schemes.GetOutlet, int64, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.Outlet) (*models.Outlet, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.Outlet) (*models.Outlet, schemes.SchemeDatabaseError)
}
