package services

import (
	"github.com/nutwreck/admin-pos-service/entities"
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type servicePaymentMethod struct {
	paymentMethod entities.EntityPaymentMethod
}

func NewServicePaymentMethod(paymentMethod entities.EntityPaymentMethod) *servicePaymentMethod {
	return &servicePaymentMethod{paymentMethod: paymentMethod}
}

/**
* =================================================
* Service Create New Master Payment Method Teritory
*==================================================
 */

func (s *servicePaymentMethod) EntityCreate(input *schemes.PaymentMethod) (*models.PaymentMethod, schemes.SchemeDatabaseError) {
	var paymentMethod schemes.PaymentMethod
	paymentMethod.MerchantID = input.MerchantID
	paymentMethod.PaymentCategoryID = input.PaymentCategoryID
	paymentMethod.AccountNumber = input.AccountNumber
	paymentMethod.Logo = input.Logo
	paymentMethod.Name = input.Name

	res, err := s.paymentMethod.EntityCreate(&paymentMethod)
	return res, err
}

/**
* ==================================================
* Service Results All Master Payment Method Teritory
*===================================================
 */

func (s *servicePaymentMethod) EntityResults(input *schemes.PaymentMethod) (*[]schemes.GetPaymentMethod, int64, schemes.SchemeDatabaseError) {
	var paymentMethod schemes.PaymentMethod
	paymentMethod.Sort = input.Sort
	paymentMethod.Page = input.Page
	paymentMethod.PerPage = input.PerPage
	paymentMethod.MerchantID = input.MerchantID
	paymentMethod.PaymentCategoryID = input.PaymentCategoryID
	paymentMethod.Name = input.Name
	paymentMethod.ID = input.ID

	res, totalData, err := s.paymentMethod.EntityResults(&paymentMethod)
	return res, totalData, err
}

/**
* ===================================================
* Service Result Master Payment Method By ID Teritory
*====================================================
 */

func (s *servicePaymentMethod) EntityResult(input *schemes.PaymentMethod) (*models.PaymentMethod, schemes.SchemeDatabaseError) {
	var paymentMethod schemes.PaymentMethod
	paymentMethod.ID = input.ID

	res, err := s.paymentMethod.EntityResult(&paymentMethod)
	return res, err
}

/**
* ===================================================
* Service Delete Master Payment Method By ID Teritory
*====================================================
 */

func (s *servicePaymentMethod) EntityDelete(input *schemes.PaymentMethod) (*models.PaymentMethod, schemes.SchemeDatabaseError) {
	var paymentMethod schemes.PaymentMethod
	paymentMethod.ID = input.ID

	res, err := s.paymentMethod.EntityDelete(&paymentMethod)
	return res, err
}

/**
* ===================================================
* Service Update Master Payment Method By ID Teritory
*====================================================
 */

func (s *servicePaymentMethod) EntityUpdate(input *schemes.PaymentMethod) (*models.PaymentMethod, schemes.SchemeDatabaseError) {
	var paymentMethod schemes.PaymentMethod
	paymentMethod.ID = input.ID
	paymentMethod.MerchantID = input.MerchantID
	paymentMethod.PaymentCategoryID = input.PaymentCategoryID
	paymentMethod.AccountNumber = input.AccountNumber
	paymentMethod.Logo = input.Logo
	paymentMethod.Name = input.Name
	paymentMethod.Active = input.Active

	res, err := s.paymentMethod.EntityUpdate(&paymentMethod)
	return res, err
}
