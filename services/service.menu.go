package services

import (
	"github.com/nutwreck/admin-pos-service/entities"
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type serviceMenu struct {
	menu entities.EntityMenu
}

func NewServiceMenu(menu entities.EntityMenu) *serviceMenu {
	return &serviceMenu{menu: menu}
}

/**
* ============================================
* Service Create New Master Menu Teritory
*=============================================
 */

func (s *serviceMenu) EntityCreate(input *schemes.Menu) (*models.Menu, schemes.SchemeDatabaseError) {
	var menu schemes.Menu
	menu.MerchantID = input.MerchantID
	menu.Name = input.Name

	res, err := s.menu.EntityCreate(&menu)
	return res, err
}

/**
* =============================================
* Service Results All Master Menu Teritory
*==============================================
 */

func (s *serviceMenu) EntityResults(input *schemes.Menu) (*[]schemes.GetMenu, int64, schemes.SchemeDatabaseError) {
	var menu schemes.Menu
	menu.Sort = input.Sort
	menu.Page = input.Page
	menu.PerPage = input.PerPage
	menu.MerchantID = input.MerchantID
	menu.Name = input.Name
	menu.ID = input.ID

	res, totalData, err := s.menu.EntityResults(&menu)
	return res, totalData, err
}

/**
* =================================================
* Service Results All Master Menu Relation Teritory
*==================================================
 */

func (s *serviceMenu) EntityResultRelations(input *schemes.Menu) (*[]schemes.GetMenuRelation, schemes.SchemeDatabaseError) {
	var menu schemes.Menu
	menu.Sort = input.Sort
	menu.MerchantID = input.MerchantID
	menu.Name = input.Name
	menu.ID = input.ID

	res, err := s.menu.EntityResultRelations(&menu)
	return res, err
}

func (s *serviceMenu) EntityGetMenu(input *schemes.Menu) (*[]schemes.GetMenu, schemes.SchemeDatabaseError) {
	var menu schemes.Menu
	menu.Sort = input.Sort
	menu.MerchantID = input.MerchantID
	menu.Name = input.Name
	menu.ID = input.ID

	res, err := s.menu.EntityGetMenu(&menu)
	return res, err
}

func (s *serviceMenu) EntityGetMenuDetail(input *schemes.MenuDetail) (*[]schemes.GetMenuDetail, schemes.SchemeDatabaseError) {
	var menuDetail schemes.MenuDetail
	menuDetail.Sort = input.Sort
	menuDetail.MerchantID = input.MerchantID
	menuDetail.MenuID = input.MenuID
	menuDetail.Name = input.Name
	menuDetail.ID = input.ID

	res, err := s.menu.EntityGetMenuDetail(&menuDetail)
	return res, err
}

func (s *serviceMenu) EntityGetMenuDetailFunction(input *schemes.MenuDetailFunction) (*[]schemes.GetMenuDetailFunction, schemes.SchemeDatabaseError) {
	var MenuDetailFunction schemes.MenuDetailFunction
	MenuDetailFunction.Sort = input.Sort
	MenuDetailFunction.MerchantID = input.MerchantID
	MenuDetailFunction.MenuDetailID = input.MenuDetailID
	MenuDetailFunction.MenuID = input.MenuID
	MenuDetailFunction.Name = input.Name
	MenuDetailFunction.ID = input.ID

	res, err := s.menu.EntityGetMenuDetailFunction(&MenuDetailFunction)
	return res, err
}

/**
* ==============================================
* Service Delete Master Menu By ID Teritory
*===============================================
 */

func (s *serviceMenu) EntityDelete(input *schemes.Menu) (*models.Menu, schemes.SchemeDatabaseError) {
	var menu schemes.Menu
	menu.ID = input.ID

	res, err := s.menu.EntityDelete(&menu)
	return res, err
}

/**
* ==============================================
* Service Update Master Menu By ID Teritory
*===============================================
 */

func (s *serviceMenu) EntityUpdate(input *schemes.Menu) (*models.Menu, schemes.SchemeDatabaseError) {
	var menu schemes.Menu
	menu.ID = input.ID
	menu.MerchantID = input.MerchantID
	menu.Name = input.Name
	menu.Active = input.Active

	res, err := s.menu.EntityUpdate(&menu)
	return res, err
}
