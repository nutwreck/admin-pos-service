package entities

import (
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type EntityPaymentMethod interface {
	EntityCreate(input *schemes.PaymentMethod) (*models.PaymentMethod, schemes.SchemeDatabaseError)
	EntityResults(input *schemes.PaymentMethod) (*[]schemes.GetPaymentMethod, int64, schemes.SchemeDatabaseError)
	EntityResult(input *schemes.PaymentMethod) (*models.PaymentMethod, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.PaymentMethod) (*models.PaymentMethod, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.PaymentMethod) (*models.PaymentMethod, schemes.SchemeDatabaseError)
}
