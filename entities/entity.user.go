package entities

import (
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type EntityUser interface {
	EntityRegister(input *schemes.SchemeUser) (*models.ModelUser, schemes.SchemeDatabaseError)
	EntityLogin(input *schemes.SchemeUser) (*models.ModelUser, schemes.SchemeDatabaseError)
	EntityGetUser(input *schemes.SchemeUser) (*models.ModelUser, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.SchemeUpdateUser) (*models.ModelUser, schemes.SchemeDatabaseError)
}
