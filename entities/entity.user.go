package entities

import (
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type EntityUser interface {
	EntityAddUser(input *schemes.User) (*models.User, schemes.SchemeDatabaseError)
	EntityLogin(input *schemes.User) (*models.User, schemes.SchemeDatabaseError)
	EntityGetUser(input *schemes.User) (*models.User, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.UpdateUser) (*models.User, schemes.SchemeDatabaseError)
	EntityGetRole(input *schemes.Role) (*models.Role, schemes.SchemeDatabaseError)
	EntityGetMerchant(input *schemes.Merchant) (*models.Merchant, schemes.SchemeDatabaseError)
	EntityGetUserOutlet(input *schemes.UserOutlet) (*[]schemes.GetUserOutlet, schemes.SchemeDatabaseError)
}
