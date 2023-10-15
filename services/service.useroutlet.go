package services

import (
	"github.com/nutwreck/admin-pos-service/entities"
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type serviceUserOutlet struct {
	userOutlet entities.EntityUserOutlet
}

func NewServiceUserOutlet(userOutlet entities.EntityUserOutlet) *serviceUserOutlet {
	return &serviceUserOutlet{userOutlet: userOutlet}
}

/**
* ==============================================
* Service Create New Master User Outlet Teritory
*===============================================
 */

func (s *serviceUserOutlet) EntityCreate(input *schemes.UserOutlet) (*models.UserOutlet, schemes.SchemeDatabaseError) {
	var userOutlet schemes.UserOutlet
	userOutlet.UserID = input.UserID
	userOutlet.OutletID = input.OutletID

	res, err := s.userOutlet.EntityCreate(&userOutlet)
	return res, err
}

/**
* ===============================================
* Service Results All Master User Outlet Teritory
*================================================
 */

func (s *serviceUserOutlet) EntityResults(input *schemes.UserOutlet) (*[]schemes.GetAllUserOutlet, int64, schemes.SchemeDatabaseError) {
	var userOutlet schemes.UserOutlet
	userOutlet.Sort = input.Sort
	userOutlet.Page = input.Page
	userOutlet.PerPage = input.PerPage
	userOutlet.MerchantID = input.MerchantID
	userOutlet.OutletID = input.OutletID
	userOutlet.UserID = input.UserID
	userOutlet.ID = input.ID

	res, totalData, err := s.userOutlet.EntityResults(&userOutlet)
	return res, totalData, err
}

/**
* ================================================
* Service Delete Master User Outlet By ID Teritory
*=================================================
 */

func (s *serviceUserOutlet) EntityDelete(input *schemes.UserOutlet) (*models.UserOutlet, schemes.SchemeDatabaseError) {
	var userOutlet schemes.UserOutlet
	userOutlet.ID = input.ID

	res, err := s.userOutlet.EntityDelete(&userOutlet)
	return res, err
}

/**
* ================================================
* Service Update Master User Outlet By ID Teritory
*=================================================
 */

func (s *serviceUserOutlet) EntityUpdate(input *schemes.UserOutlet) (*models.UserOutlet, schemes.SchemeDatabaseError) {
	var userOutlet schemes.UserOutlet
	userOutlet.ID = input.ID
	userOutlet.OutletID = input.OutletID
	userOutlet.UserID = input.UserID
	userOutlet.Active = input.Active

	res, err := s.userOutlet.EntityUpdate(&userOutlet)
	return res, err
}
