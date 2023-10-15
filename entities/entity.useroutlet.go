package entities

import (
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type EntityUserOutlet interface {
	EntityCreate(input *schemes.UserOutlet) (*models.UserOutlet, schemes.SchemeDatabaseError)
	EntityResults(input *schemes.UserOutlet) (*[]schemes.GetAllUserOutlet, int64, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.UserOutlet) (*models.UserOutlet, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.UserOutlet) (*models.UserOutlet, schemes.SchemeDatabaseError)
}
