package services

import (
	"github.com/nutwreck/admin-pos-service/entities"
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type serviceUnitOfMeasurementType struct {
	uomType entities.EntityUnitOfMeasurementType
}

func NewServiceUnitOfMeasurementType(uomType entities.EntityUnitOfMeasurementType) *serviceUnitOfMeasurementType {
	return &serviceUnitOfMeasurementType{uomType: uomType}
}

/**
* ============================================
* Service Create New Master UOM Type Teritory
*=============================================
 */

func (s *serviceUnitOfMeasurementType) EntityCreate(input *schemes.UnitOfMeasurementType) (*models.UnitOfMeasurementType, schemes.SchemeDatabaseError) {
	var uomType schemes.UnitOfMeasurementType
	uomType.MerchantID = input.MerchantID
	uomType.Name = input.Name

	res, err := s.uomType.EntityCreate(&uomType)
	return res, err
}

/**
* =============================================
* Service Results All Master UOM Type Teritory
*==============================================
 */

func (s *serviceUnitOfMeasurementType) EntityResults(input *schemes.UnitOfMeasurementType) (*[]schemes.GetUnitOfMeasurementType, int64, schemes.SchemeDatabaseError) {
	var uomType schemes.UnitOfMeasurementType
	uomType.Sort = input.Sort
	uomType.Page = input.Page
	uomType.PerPage = input.PerPage
	uomType.MerchantID = input.MerchantID
	uomType.Name = input.Name
	uomType.ID = input.ID

	res, totalData, err := s.uomType.EntityResults(&uomType)
	return res, totalData, err
}

/**
* ==============================================
* Service Delete Master UOM Type By ID Teritory
*===============================================
 */

func (s *serviceUnitOfMeasurementType) EntityDelete(input *schemes.UnitOfMeasurementType) (*models.UnitOfMeasurementType, schemes.SchemeDatabaseError) {
	var uomType schemes.UnitOfMeasurementType
	uomType.ID = input.ID

	res, err := s.uomType.EntityDelete(&uomType)
	return res, err
}

/**
* ==============================================
* Service Update Master UOM Type By ID Teritory
*===============================================
 */

func (s *serviceUnitOfMeasurementType) EntityUpdate(input *schemes.UnitOfMeasurementType) (*models.UnitOfMeasurementType, schemes.SchemeDatabaseError) {
	var uomType schemes.UnitOfMeasurementType
	uomType.ID = input.ID
	uomType.MerchantID = input.MerchantID
	uomType.Name = input.Name
	uomType.Active = input.Active

	res, err := s.uomType.EntityUpdate(&uomType)
	return res, err
}
