package services

import (
	"github.com/nutwreck/admin-pos-service/entities"
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type serviceOutlet struct {
	outlet entities.EntityOutlet
}

func NewServiceOutlet(outlet entities.EntityOutlet) *serviceOutlet {
	return &serviceOutlet{outlet: outlet}
}

/**
* ==========================================
* Service Create New Outlet Teritory
*===========================================
 */

func (s *serviceOutlet) EntityCreate(input *schemes.Outlet) (*models.Outlet, schemes.SchemeDatabaseError) {
	var outlet schemes.Outlet
	outlet.Name = input.Name
	outlet.Phone = input.Phone
	outlet.Address = input.Address
	outlet.MerchantID = input.MerchantID
	outlet.Description = input.Description

	res, err := s.outlet.EntityCreate(&outlet)
	return res, err
}

/**
* ==========================================
* Service Result All Outlet Teritory
*===========================================
 */

func (s *serviceOutlet) EntityResults(input *schemes.Outlet) (*[]schemes.GetOutlet, int64, schemes.SchemeDatabaseError) {
	var outlet schemes.Outlet
	outlet.Sort = input.Sort
	outlet.Page = input.Page
	outlet.PerPage = input.PerPage
	outlet.Name = input.Name
	outlet.ID = input.ID

	res, totalData, err := s.outlet.EntityResults(&outlet)
	return res, totalData, err
}

/**
* ==========================================
* Service Delete Outlet By ID Teritory
*===========================================
 */

func (s *serviceOutlet) EntityDelete(input *schemes.Outlet) (*models.Outlet, schemes.SchemeDatabaseError) {
	var outlet schemes.Outlet
	outlet.ID = input.ID

	res, err := s.outlet.EntityDelete(&outlet)
	return res, err
}

/**
* ==========================================
* Service Update Outlet By ID Teritory
*===========================================
 */

func (s *serviceOutlet) EntityUpdate(input *schemes.Outlet) (*models.Outlet, schemes.SchemeDatabaseError) {
	var outlet schemes.Outlet
	outlet.ID = input.ID
	outlet.Name = input.Name
	outlet.Phone = input.Phone
	outlet.Address = input.Address
	outlet.MerchantID = input.MerchantID
	outlet.Description = input.Description
	outlet.Active = input.Active

	res, err := s.outlet.EntityUpdate(&outlet)
	return res, err
}
