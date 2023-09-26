package services

import (
	"github.com/nutwreck/admin-pos-service/entities"
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type serviceMenuDetail struct {
	menuDetail entities.EntityMenuDetail
}

func NewServiceMenuDetail(menuDetail entities.EntityMenuDetail) *serviceMenuDetail {
	return &serviceMenuDetail{menuDetail: menuDetail}
}

/**
* ==============================================
* Service Create New Master Menu Detail Teritory
*===============================================
 */

func (s *serviceMenuDetail) EntityCreate(input *schemes.MenuDetail) (*models.MenuDetail, schemes.SchemeDatabaseError) {
	var menuDetail schemes.MenuDetail
	menuDetail.Name = input.Name
	menuDetail.MenuID = input.MenuID
	menuDetail.Link = input.Link
	menuDetail.Image = input.Image
	menuDetail.Icon = input.Icon

	res, err := s.menuDetail.EntityCreate(&menuDetail)
	return res, err
}

/**
* ===============================================
* Service Results All Master Menu Detail Teritory
*================================================
 */

func (s *serviceMenuDetail) EntityResults(input *schemes.MenuDetail) (*[]schemes.GetMenuDetail, int64, schemes.SchemeDatabaseError) {
	var menuDetail schemes.MenuDetail
	menuDetail.Sort = input.Sort
	menuDetail.Page = input.Page
	menuDetail.PerPage = input.PerPage
	menuDetail.Name = input.Name
	menuDetail.ID = input.ID

	res, totalData, err := s.menuDetail.EntityResults(&menuDetail)
	return res, totalData, err
}

/**
* ================================================
* Service Delete Master Menu Detail By ID Teritory
*=================================================
 */

func (s *serviceMenuDetail) EntityDelete(input *schemes.MenuDetail) (*models.MenuDetail, schemes.SchemeDatabaseError) {
	var menuDetail schemes.MenuDetail
	menuDetail.ID = input.ID

	res, err := s.menuDetail.EntityDelete(&menuDetail)
	return res, err
}

/**
* ================================================
* Service Update Master Menu Detail By ID Teritory
*=================================================
 */

func (s *serviceMenuDetail) EntityUpdate(input *schemes.MenuDetail) (*models.MenuDetail, schemes.SchemeDatabaseError) {
	var menuDetail schemes.MenuDetail
	menuDetail.ID = input.ID
	menuDetail.Name = input.Name
	menuDetail.MenuID = input.MenuID
	menuDetail.Link = input.Link
	menuDetail.Image = input.Image
	menuDetail.Icon = input.Icon
	menuDetail.Active = input.Active

	res, err := s.menuDetail.EntityUpdate(&menuDetail)
	return res, err
}

/**
* ==========================================
* Service Result Master Menu Detail Teritory
*===========================================
 */

func (s *serviceMenuDetail) EntityResult(input *schemes.MenuDetail) (*models.MenuDetail, schemes.SchemeDatabaseError) {
	var menuDetail schemes.MenuDetail
	menuDetail.ID = input.ID

	res, err := s.menuDetail.EntityResult(&menuDetail)
	return res, err
}
