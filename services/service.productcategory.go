package services

import (
	"github.com/nutwreck/admin-pos-service/entities"
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type serviceProductCategory struct {
	productcategory entities.EntityProductCategory
}

func NewServiceProductCategory(productcategory entities.EntityProductCategory) *serviceProductCategory {
	return &serviceProductCategory{productcategory: productcategory}
}

/**
* ========================================================
* Service Create New Master Product Category Type Teritory
*=========================================================
 */

func (s *serviceProductCategory) EntityCreate(input *schemes.ProductCategory) (*models.ProductCategory, schemes.SchemeDatabaseError) {
	var productcategory schemes.ProductCategory
	productcategory.MerchantID = input.MerchantID
	productcategory.OutletID = input.OutletID
	productcategory.Name = input.Name

	res, err := s.productcategory.EntityCreate(&productcategory)
	return res, err
}

/**
* =========================================================
* Service Results All Master Product Category Type Teritory
*==========================================================
 */

func (s *serviceProductCategory) EntityResults(input *schemes.ProductCategory) (*[]schemes.GetProductCategory, int64, schemes.SchemeDatabaseError) {
	var productcategory schemes.ProductCategory
	productcategory.Sort = input.Sort
	productcategory.Page = input.Page
	productcategory.PerPage = input.PerPage
	productcategory.MerchantID = input.MerchantID
	productcategory.OutletID = input.OutletID
	productcategory.Name = input.Name
	productcategory.ID = input.ID

	res, totalData, err := s.productcategory.EntityResults(&productcategory)
	return res, totalData, err
}

/**
* ==========================================================
* Service Delete Master Product Category Type By ID Teritory
*===========================================================
 */

func (s *serviceProductCategory) EntityDelete(input *schemes.ProductCategory) (*models.ProductCategory, schemes.SchemeDatabaseError) {
	var productcategory schemes.ProductCategory
	productcategory.ID = input.ID

	res, err := s.productcategory.EntityDelete(&productcategory)
	return res, err
}

/**
* ==========================================================
* Service Update Master Product Category Type By ID Teritory
*===========================================================
 */

func (s *serviceProductCategory) EntityUpdate(input *schemes.ProductCategory) (*models.ProductCategory, schemes.SchemeDatabaseError) {
	var productcategory schemes.ProductCategory
	productcategory.ID = input.ID
	productcategory.MerchantID = input.MerchantID
	productcategory.OutletID = input.OutletID
	productcategory.Name = input.Name
	productcategory.Active = input.Active

	res, err := s.productcategory.EntityUpdate(&productcategory)
	return res, err
}
