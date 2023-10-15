package entities

import (
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type EntityPaymentCategory interface {
	EntityCreate(input *schemes.PaymentCategory) (*models.PaymentCategory, schemes.SchemeDatabaseError)
	EntityResults(input *schemes.PaymentCategory) (*[]schemes.GetPaymentCategory, int64, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.PaymentCategory) (*models.PaymentCategory, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.PaymentCategory) (*models.PaymentCategory, schemes.SchemeDatabaseError)
}
