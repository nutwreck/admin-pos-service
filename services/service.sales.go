package services

import (
	"github.com/nutwreck/admin-pos-service/entities"
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type serviceSales struct {
	sales entities.EntitySales
}

func NewServiceSales(sales entities.EntitySales) *serviceSales {
	return &serviceSales{sales: sales}
}

/**
* ==========================================
* Service Create New Sales Teritory
*===========================================
 */

func (s *serviceSales) EntityCreate(input *schemes.Sales) (*models.Sales, schemes.SchemeDatabaseError) {
	var sales schemes.Sales
	sales.Name = input.Name
	sales.Phone = input.Phone
	sales.Address = input.Address
	sales.Description = input.Description
	sales.MerchantID = input.MerchantID
	sales.OutletID = input.OutletID

	res, err := s.sales.EntityCreate(&sales)
	return res, err
}

/**
* ==========================================
* Service Results All Sales Teritory
*===========================================
 */

func (s *serviceSales) EntityResults(input *schemes.Sales) (*[]schemes.GetSales, int64, schemes.SchemeDatabaseError) {
	var sales schemes.Sales
	sales.Sort = input.Sort
	sales.Page = input.Page
	sales.PerPage = input.PerPage
	sales.MerchantID = input.MerchantID
	sales.OutletID = input.OutletID
	sales.Name = input.Name
	sales.ID = input.ID

	res, totalData, err := s.sales.EntityResults(&sales)
	return res, totalData, err
}

/**
* ==========================================
* Service Result Sales By ID Teritory
*===========================================
 */

func (s *serviceSales) EntityResult(input *schemes.Sales) (*models.Sales, schemes.SchemeDatabaseError) {
	var sales schemes.Sales
	sales.ID = input.ID

	res, err := s.sales.EntityResult(&sales)
	return res, err
}

/**
* ==========================================
* Service Delete Sales By ID Teritory
*===========================================
 */

func (s *serviceSales) EntityDelete(input *schemes.Sales) (*models.Sales, schemes.SchemeDatabaseError) {
	var sales schemes.Sales
	sales.ID = input.ID

	res, err := s.sales.EntityDelete(&sales)
	return res, err
}

/**
* ==========================================
* Service Update Sales By ID Teritory
*===========================================
 */

func (s *serviceSales) EntityUpdate(input *schemes.Sales) (*models.Sales, schemes.SchemeDatabaseError) {
	var sales schemes.Sales
	sales.ID = input.ID
	sales.Name = input.Name
	sales.Phone = input.Phone
	sales.Address = input.Address
	sales.Description = input.Description
	sales.MerchantID = input.MerchantID
	sales.OutletID = input.OutletID
	sales.Active = input.Active

	res, err := s.sales.EntityUpdate(&sales)
	return res, err
}
