package services

import (
	"github.com/nutwreck/admin-pos-service/entities"
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type serviceUnitOfMeasurement struct {
	uom entities.EntityUnitOfMeasurement
}

func NewServiceUnitOfMeasurement(uom entities.EntityUnitOfMeasurement) *serviceUnitOfMeasurement {
	return &serviceUnitOfMeasurement{uom: uom}
}

/**
* ============================================
* Service Create New Master UOM Teritory
*=============================================
 */

func (s *serviceUnitOfMeasurement) EntityCreate(input *schemes.UnitOfMeasurement) (*models.UnitOfMeasurement, schemes.SchemeDatabaseError) {
	var uom schemes.UnitOfMeasurement
	uom.MerchantID = input.MerchantID
	uom.UOMTypeID = input.UOMTypeID
	uom.Symbol = input.Symbol
	uom.ConversionFactor = input.ConversionFactor
	uom.Name = input.Name

	res, err := s.uom.EntityCreate(&uom)
	return res, err
}

/**
* =============================================
* Service Results All Master UOM Teritory
*==============================================
 */

func (s *serviceUnitOfMeasurement) EntityResults(input *schemes.UnitOfMeasurement) (*[]schemes.GetUnitOfMeasurement, int64, schemes.SchemeDatabaseError) {
	var uom schemes.UnitOfMeasurement
	uom.Sort = input.Sort
	uom.Page = input.Page
	uom.PerPage = input.PerPage
	uom.MerchantID = input.MerchantID
	uom.UOMTypeID = input.UOMTypeID
	uom.Name = input.Name
	uom.ID = input.ID

	res, totalData, err := s.uom.EntityResults(&uom)
	return res, totalData, err
}

/**
* ==============================================
* Service Delete Master UOM By ID Teritory
*===============================================
 */

func (s *serviceUnitOfMeasurement) EntityDelete(input *schemes.UnitOfMeasurement) (*models.UnitOfMeasurement, schemes.SchemeDatabaseError) {
	var uom schemes.UnitOfMeasurement
	uom.ID = input.ID

	res, err := s.uom.EntityDelete(&uom)
	return res, err
}

/**
* ==============================================
* Service Update Master UOM By ID Teritory
*===============================================
 */

func (s *serviceUnitOfMeasurement) EntityUpdate(input *schemes.UnitOfMeasurement) (*models.UnitOfMeasurement, schemes.SchemeDatabaseError) {
	var uom schemes.UnitOfMeasurement
	uom.ID = input.ID
	uom.MerchantID = input.MerchantID
	uom.UOMTypeID = input.UOMTypeID
	uom.Symbol = input.Symbol
	uom.ConversionFactor = input.ConversionFactor
	uom.Name = input.Name
	uom.Active = input.Active

	res, err := s.uom.EntityUpdate(&uom)
	return res, err
}
