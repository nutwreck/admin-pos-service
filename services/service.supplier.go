package services

import (
	"github.com/nutwreck/admin-pos-service/entities"
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type serviceSupplier struct {
	supplier entities.EntitySupplier
}

func NewServiceSupplier(supplier entities.EntitySupplier) *serviceSupplier {
	return &serviceSupplier{supplier: supplier}
}

/**
* ==========================================
* Service Create New Supplier Teritory
*===========================================
 */

func (s *serviceSupplier) EntityCreate(input *schemes.Supplier) (*models.Supplier, schemes.SchemeDatabaseError) {
	var supplier schemes.Supplier
	supplier.Name = input.Name
	supplier.Phone = input.Phone
	supplier.Address = input.Address
	supplier.Description = input.Description
	supplier.MerchantID = input.MerchantID
	supplier.OutletID = input.OutletID

	res, err := s.supplier.EntityCreate(&supplier)
	return res, err
}

/**
* ==========================================
* Service Results All Supplier Teritory
*===========================================
 */

func (s *serviceSupplier) EntityResults(input *schemes.Supplier) (*[]schemes.GetSupplier, int64, schemes.SchemeDatabaseError) {
	var supplier schemes.Supplier
	supplier.Sort = input.Sort
	supplier.Page = input.Page
	supplier.PerPage = input.PerPage
	supplier.Name = input.Name
	supplier.ID = input.ID

	res, totalData, err := s.supplier.EntityResults(&supplier)
	return res, totalData, err
}

/**
* ==========================================
* Service Result Supplier By ID Teritory
*===========================================
 */

func (s *serviceSupplier) EntityResult(input *schemes.Supplier) (*models.Supplier, schemes.SchemeDatabaseError) {
	var supplier schemes.Supplier
	supplier.ID = input.ID

	res, err := s.supplier.EntityResult(&supplier)
	return res, err
}

/**
* ==========================================
* Service Delete Supplier By ID Teritory
*===========================================
 */

func (s *serviceSupplier) EntityDelete(input *schemes.Supplier) (*models.Supplier, schemes.SchemeDatabaseError) {
	var supplier schemes.Supplier
	supplier.ID = input.ID

	res, err := s.supplier.EntityDelete(&supplier)
	return res, err
}

/**
* ==========================================
* Service Update Supplier By ID Teritory
*===========================================
 */

func (s *serviceSupplier) EntityUpdate(input *schemes.Supplier) (*models.Supplier, schemes.SchemeDatabaseError) {
	var supplier schemes.Supplier
	supplier.ID = input.ID
	supplier.Name = input.Name
	supplier.Phone = input.Phone
	supplier.Address = input.Address
	supplier.Description = input.Description
	supplier.MerchantID = input.MerchantID
	supplier.OutletID = input.OutletID
	supplier.Active = input.Active

	res, err := s.supplier.EntityUpdate(&supplier)
	return res, err
}
