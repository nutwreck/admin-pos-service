package services

import (
	"github.com/nutwreck/admin-pos-service/entities"
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type servicePaymentCategory struct {
	paymentCategory entities.EntityPaymentCategory
}

func NewServicePaymentCategory(paymentCategory entities.EntityPaymentCategory) *servicePaymentCategory {
	return &servicePaymentCategory{paymentCategory: paymentCategory}
}

/**
* ===================================================
* Service Create New Master Payment Category Teritory
*====================================================
 */

func (s *servicePaymentCategory) EntityCreate(input *schemes.PaymentCategory) (*models.PaymentCategory, schemes.SchemeDatabaseError) {
	var paymentCategory schemes.PaymentCategory
	paymentCategory.MerchantID = input.MerchantID
	paymentCategory.Name = input.Name

	res, err := s.paymentCategory.EntityCreate(&paymentCategory)
	return res, err
}

/**
* ====================================================
* Service Results All Master Payment Category Teritory
*=====================================================
 */

func (s *servicePaymentCategory) EntityResults(input *schemes.PaymentCategory) (*[]schemes.GetPaymentCategory, int64, schemes.SchemeDatabaseError) {
	var paymentCategory schemes.PaymentCategory
	paymentCategory.Sort = input.Sort
	paymentCategory.Page = input.Page
	paymentCategory.PerPage = input.PerPage
	paymentCategory.MerchantID = input.MerchantID
	paymentCategory.Name = input.Name
	paymentCategory.ID = input.ID

	res, totalData, err := s.paymentCategory.EntityResults(&paymentCategory)
	return res, totalData, err
}

/**
* =====================================================
* Service Delete Master Payment Category By ID Teritory
*======================================================
 */

func (s *servicePaymentCategory) EntityDelete(input *schemes.PaymentCategory) (*models.PaymentCategory, schemes.SchemeDatabaseError) {
	var paymentCategory schemes.PaymentCategory
	paymentCategory.ID = input.ID

	res, err := s.paymentCategory.EntityDelete(&paymentCategory)
	return res, err
}

/**
* =====================================================
* Service Update Master Payment Category By ID Teritory
*======================================================
 */

func (s *servicePaymentCategory) EntityUpdate(input *schemes.PaymentCategory) (*models.PaymentCategory, schemes.SchemeDatabaseError) {
	var paymentCategory schemes.PaymentCategory
	paymentCategory.ID = input.ID
	paymentCategory.MerchantID = input.MerchantID
	paymentCategory.Name = input.Name
	paymentCategory.Active = input.Active

	res, err := s.paymentCategory.EntityUpdate(&paymentCategory)
	return res, err
}
