package services

import (
	"github.com/nutwreck/admin-pos-service/entities"
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type serviceProductCategorySub struct {
	productcategorysub entities.EntityProductCategorySub
}

func NewServiceProductCategorySub(productcategorysub entities.EntityProductCategorySub) *serviceProductCategorySub {
	return &serviceProductCategorySub{productcategorysub: productcategorysub}
}

/**
* ========================================================
* Service Create New Master Product Category Sub Teritory
*=========================================================
 */

func (s *serviceProductCategorySub) EntityCreate(input *schemes.ProductCategorySub) (*models.ProductCategorySub, schemes.SchemeDatabaseError) {
	var productcategorysub schemes.ProductCategorySub
	productcategorysub.MerchantID = input.MerchantID
	productcategorysub.OutletID = input.OutletID
	productcategorysub.ProductCategoryID = input.ProductCategoryID
	productcategorysub.Name = input.Name

	res, err := s.productcategorysub.EntityCreate(&productcategorysub)
	return res, err
}

/**
* =========================================================
* Service Results All Master Product Category Sub Teritory
*==========================================================
 */

func (s *serviceProductCategorySub) EntityResults(input *schemes.ProductCategorySub) (*[]schemes.GetProductCategorySub, int64, schemes.SchemeDatabaseError) {
	var productcategorysub schemes.ProductCategorySub
	productcategorysub.Sort = input.Sort
	productcategorysub.Page = input.Page
	productcategorysub.PerPage = input.PerPage
	productcategorysub.MerchantID = input.MerchantID
	productcategorysub.OutletID = input.OutletID
	productcategorysub.ProductCategoryID = input.ProductCategoryID
	productcategorysub.Name = input.Name
	productcategorysub.ID = input.ID

	res, totalData, err := s.productcategorysub.EntityResults(&productcategorysub)
	return res, totalData, err
}

/**
* ==========================================================
* Service Delete Master Product Category Sub By ID Teritory
*===========================================================
 */

func (s *serviceProductCategorySub) EntityDelete(input *schemes.ProductCategorySub) (*models.ProductCategorySub, schemes.SchemeDatabaseError) {
	var productcategorysub schemes.ProductCategorySub
	productcategorysub.ID = input.ID

	res, err := s.productcategorysub.EntityDelete(&productcategorysub)
	return res, err
}

/**
* ==========================================================
* Service Update Master Product Category Sub By ID Teritory
*===========================================================
 */

func (s *serviceProductCategorySub) EntityUpdate(input *schemes.ProductCategorySub) (*models.ProductCategorySub, schemes.SchemeDatabaseError) {
	var productcategorysub schemes.ProductCategorySub
	productcategorysub.ID = input.ID
	productcategorysub.MerchantID = input.MerchantID
	productcategorysub.OutletID = input.OutletID
	productcategorysub.ProductCategoryID = input.ProductCategoryID
	productcategorysub.Name = input.Name
	productcategorysub.Active = input.Active

	res, err := s.productcategorysub.EntityUpdate(&productcategorysub)
	return res, err
}
