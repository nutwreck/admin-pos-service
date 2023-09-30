package entities

import (
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type EntityMerchant interface {
	EntityCreate(input *schemes.Merchant) (*models.Merchant, schemes.SchemeDatabaseError)
	EntityResults(input *schemes.Merchant) (*[]schemes.GetMerchant, int64, schemes.SchemeDatabaseError)
	EntityResult(input *schemes.Merchant) (*models.Merchant, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.Merchant) (*models.Merchant, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.Merchant) (*models.Merchant, schemes.SchemeDatabaseError)
}
