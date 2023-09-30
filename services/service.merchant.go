package services

import (
	"github.com/nutwreck/admin-pos-service/entities"
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type serviceMerchant struct {
	merchant entities.EntityMerchant
}

func NewServiceMerchant(merchant entities.EntityMerchant) *serviceMerchant {
	return &serviceMerchant{merchant: merchant}
}

/**
* ==========================================
* Service Create New Merchant Teritory
*===========================================
 */

func (s *serviceMerchant) EntityCreate(input *schemes.Merchant) (*models.Merchant, schemes.SchemeDatabaseError) {
	var merchant schemes.Merchant
	merchant.Name = input.Name
	merchant.Phone = input.Phone
	merchant.Address = input.Address
	merchant.Logo = input.Logo
	merchant.Description = input.Description

	res, err := s.merchant.EntityCreate(&merchant)
	return res, err
}

/**
* ==========================================
* Service Results All MerchantTeritory
*===========================================
 */

func (s *serviceMerchant) EntityResults(input *schemes.Merchant) (*[]schemes.GetMerchant, int64, schemes.SchemeDatabaseError) {
	var merchant schemes.Merchant
	merchant.Sort = input.Sort
	merchant.Page = input.Page
	merchant.PerPage = input.PerPage
	merchant.Name = input.Name
	merchant.ID = input.ID

	res, totalData, err := s.merchant.EntityResults(&merchant)
	return res, totalData, err
}

/**
* ==========================================
* Service Result Merchant By ID Teritory
*===========================================
 */

func (s *serviceMerchant) EntityResult(input *schemes.Merchant) (*models.Merchant, schemes.SchemeDatabaseError) {
	var merchant schemes.Merchant
	merchant.ID = input.ID

	res, err := s.merchant.EntityResult(&merchant)
	return res, err
}

/**
* ==========================================
* Service Delete Merchant By ID Teritory
*===========================================
 */

func (s *serviceMerchant) EntityDelete(input *schemes.Merchant) (*models.Merchant, schemes.SchemeDatabaseError) {
	var merchant schemes.Merchant
	merchant.ID = input.ID

	res, err := s.merchant.EntityDelete(&merchant)
	return res, err
}

/**
* ==========================================
* Service Update Merchant By ID Teritory
*===========================================
 */

func (s *serviceMerchant) EntityUpdate(input *schemes.Merchant) (*models.Merchant, schemes.SchemeDatabaseError) {
	var merchant schemes.Merchant
	merchant.ID = input.ID
	merchant.Name = input.Name
	merchant.Phone = input.Phone
	merchant.Address = input.Address
	merchant.Logo = input.Logo
	merchant.Description = input.Description
	merchant.Active = input.Active

	res, err := s.merchant.EntityUpdate(&merchant)
	return res, err
}
