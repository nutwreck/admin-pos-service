package services

import (
	"github.com/nutwreck/admin-pos-service/entities"
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type serviceMenuDetailFunction struct {
	menuDetailFunction entities.EntityMenuDetailFunction
}

func NewServiceMenuDetailFunction(menuDetailFunction entities.EntityMenuDetailFunction) *serviceMenuDetailFunction {
	return &serviceMenuDetailFunction{menuDetailFunction: menuDetailFunction}
}

/**
* =======================================================
* Service Create New Master Menu Detail Function Teritory
*========================================================
 */

func (s *serviceMenuDetailFunction) EntityCreate(inputs *[]schemes.MenuDetailFunction) (*models.MenuDetailFunction, schemes.SchemeDatabaseError) {
	var createdMenuDetailFunctions []schemes.MenuDetailFunction

	// Loop through each input in the batch
	for _, input := range *inputs {
		var menuDetailFunction schemes.MenuDetailFunction
		menuDetailFunction.MerchantID = input.MerchantID
		menuDetailFunction.Name = input.Name
		menuDetailFunction.MenuID = input.MenuID
		menuDetailFunction.MenuDetailID = input.MenuDetailID
		menuDetailFunction.Link = input.Link

		// Collect the created MenuDetailFunction objects
		createdMenuDetailFunctions = append(createdMenuDetailFunctions, menuDetailFunction)
	}

	res, err := s.menuDetailFunction.EntityCreate(&createdMenuDetailFunctions)
	return res, err
}

/**
* ========================================================
* Service Results All Master Menu Detail Function Teritory
*=========================================================
 */

func (s *serviceMenuDetailFunction) EntityResults(input *schemes.MenuDetailFunction) (*[]schemes.GetMenuDetailFunction, int64, schemes.SchemeDatabaseError) {
	var menuDetailFunction schemes.MenuDetailFunction
	menuDetailFunction.Sort = input.Sort
	menuDetailFunction.Page = input.Page
	menuDetailFunction.PerPage = input.PerPage
	menuDetailFunction.MerchantID = input.MerchantID
	menuDetailFunction.Name = input.Name
	menuDetailFunction.ID = input.ID

	res, totalData, err := s.menuDetailFunction.EntityResults(&menuDetailFunction)
	return res, totalData, err
}

/**
* =========================================================
* Service Delete Master Menu Detail Function By ID Teritory
*==========================================================
 */

func (s *serviceMenuDetailFunction) EntityDelete(input *schemes.MenuDetailFunction) (*models.MenuDetailFunction, schemes.SchemeDatabaseError) {
	var menuDetailFunction schemes.MenuDetailFunction
	menuDetailFunction.ID = input.ID

	res, err := s.menuDetailFunction.EntityDelete(&menuDetailFunction)
	return res, err
}

/**
* ================================================
* Service Update Master Menu Detail Function By ID Teritory
*=================================================
 */

func (s *serviceMenuDetailFunction) EntityUpdate(input *schemes.MenuDetailFunction) (*models.MenuDetailFunction, schemes.SchemeDatabaseError) {
	var menuDetailFunction schemes.MenuDetailFunction
	menuDetailFunction.ID = input.ID
	menuDetailFunction.MerchantID = input.MerchantID
	menuDetailFunction.Name = input.Name
	menuDetailFunction.MenuID = input.MenuID
	menuDetailFunction.MenuDetailID = input.MenuDetailID
	menuDetailFunction.Link = input.Link
	menuDetailFunction.Active = input.Active

	res, err := s.menuDetailFunction.EntityUpdate(&menuDetailFunction)
	return res, err
}

/**
* ===================================================
* Service Result Master Menu Detail Function Teritory
*====================================================
 */

func (s *serviceMenuDetailFunction) EntityResult(input *schemes.MenuDetailFunction) (*models.MenuDetailFunction, schemes.SchemeDatabaseError) {
	var menuDetailFunction schemes.MenuDetailFunction
	menuDetailFunction.ID = input.ID

	res, err := s.menuDetailFunction.EntityResult(&menuDetailFunction)
	return res, err
}
