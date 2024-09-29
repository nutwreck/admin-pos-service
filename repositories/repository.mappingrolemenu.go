package repositories

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/nutwreck/admin-pos-service/constants"
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
	"gorm.io/gorm"
)

type repositoryMappingRoleMenu struct {
	db *gorm.DB
}

func NewRepositoryMappingRoleMenu(db *gorm.DB) *repositoryMappingRoleMenu {
	return &repositoryMappingRoleMenu{db: db}
}

/**
* ======================================================
* Repository Create New Mapping Role Menu Teritory
*=======================================================
 */

func (r *repositoryMappingRoleMenu) EntityCreate(input *schemes.MappingRoleMenu) (*models.MappingRoleMenu, schemes.SchemeDatabaseError) {
	var mappingRoleMenu models.MappingRoleMenu
	mappingRoleMenu.MerchantID = input.MerchantID
	mappingRoleMenu.RoleID = input.RoleID
	mappingRoleMenu.MenuID = input.MenuID
	mappingRoleMenu.MenuDetailID = input.MenuDetailID
	mappingRoleMenu.MenuDetailFunctionID = input.MenuDetailFunctionID
	mappingRoleMenu.Name = input.Name

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&mappingRoleMenu)

	checkData := db.Debug().Where("merchant_id = ? AND name = ?", mappingRoleMenu.MerchantID, mappingRoleMenu.Name).First(&mappingRoleMenu)

	if checkData.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &mappingRoleMenu, <-err
	}

	addData := db.Debug().Create(&mappingRoleMenu).Commit()

	if addData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_02",
		}
		return &mappingRoleMenu, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &mappingRoleMenu, <-err
}

/**
* =======================================================
* Repository Results All Mapping Role Menu Teritory
*========================================================
 */

func (r *repositoryMappingRoleMenu) EntityResults(input *schemes.MappingRoleMenu) (*[]schemes.GetMappingRoleMenu, int64, schemes.SchemeDatabaseError) {
	var (
		mappingRoleMenu []models.MappingRoleMenu
		result          []schemes.GetMappingRoleMenu
		countData       schemes.CountData
		args            []interface{}
		totalData       int64
		sortData        string = "mappingrolemenu.created_at DESC"
		queryCountData  string = constants.EMPTY_VALUE
		queryData       string = constants.EMPTY_VALUE
		queryAdditional string = constants.EMPTY_VALUE
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&mappingRoleMenu)

	if input.Sort != constants.EMPTY_VALUE {
		unScape, _ := url.QueryUnescape(input.Sort)
		sortData = strings.Replace(unScape, "'", constants.EMPTY_VALUE, -1)
	}

	offset := int((input.Page - 1) * input.PerPage)

	//Untuk mengambil jumlah data tanpa limit
	queryCountData = `
		SELECT
			COUNT(mappingrolemenu.*) AS count_data
		FROM master.mapping_role_menus AS mappingrolemenu
	`

	//Untuk mengambil detail data
	queryData = `
		SELECT
			mappingrolemenu.id,
			mappingrolemenu.name,
			mappingrolemenu.active,
			mappingrolemenu.created_at,
			merchant.id AS merchant_id,
			merchant.name AS merchant_name,
			role.id AS role_id,
			role.name AS role_name,
			menu.id AS menu_id,
			menu.name AS menu_name,
			menudetail.id AS menu_detail_id,
			menudetail.name AS menu_detail_name,
			menudetailfunction.id AS menu_detail_function_id,
			menudetailfunction.name AS menu_detail_function_name
		FROM master.mapping_role_menus AS mappingrolemenu
	`

	queryAdditional = `
		JOIN master.merchants AS merchant ON mappingrolemenu.merchant_id = merchant.id AND merchant.active = true
		JOIN master.roles AS role ON mappingrolemenu.role_id = role.id
		JOIN master.menus AS menu ON mappingrolemenu.menu_id = menu.id
		JOIN master.menu_details AS menudetail ON mappingrolemenu.menu_detail_id = menudetail.id
		JOIN master.menu_detail_functions AS menudetailfunction ON mappingrolemenu.menu_detail_function_id = menudetailfunction.id
	`

	queryAdditional += ` WHERE TRUE`

	if input.MerchantID != constants.EMPTY_VALUE {
		queryAdditional += ` AND mappingrolemenu.merchant_id = ?`
		args = append(args, input.MerchantID)
	}

	if input.Name != constants.EMPTY_VALUE {
		queryAdditional += ` AND mappingrolemenu.name LIKE ?`
		args = append(args, "%"+input.Name+"%")
	}

	if input.ID != constants.EMPTY_VALUE {
		queryAdditional += ` AND mappingrolemenu.id = ?`
		args = append(args, input.ID)
	}

	//Eksekusi query ambil jumlah data tanpa limit
	db.Raw(queryCountData+queryAdditional, args...).Scan(&countData)

	queryAdditional += ` ORDER BY ` + sortData

	if input.Page != constants.EMPTY_NUMBER || input.PerPage != constants.EMPTY_NUMBER {
		queryAdditional += ` LIMIT ?`
		args = append(args, int(input.PerPage))

		queryAdditional += ` OFFSET ?`
		args = append(args, offset)
	}

	getDatas := db.Raw(queryData+queryAdditional, args...).Scan(&result)

	if getDatas.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &result, totalData, <-err
	}

	// Menghitung total data yang diambil
	totalData = countData.CountData

	err <- schemes.SchemeDatabaseError{}
	return &result, totalData, <-err
}

/**
* ========================================================
* Repository Delete Mapping Role Menu By ID Teritory
*=========================================================
 */

func (r *repositoryMappingRoleMenu) EntityDelete(input *schemes.MappingRoleMenu) (*models.MappingRoleMenu, schemes.SchemeDatabaseError) {
	var mappingRoleMenu models.MappingRoleMenu
	mappingRoleMenu.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&mappingRoleMenu)

	checkId := db.Debug().First(&mappingRoleMenu)

	if checkId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &mappingRoleMenu, <-err
	}

	deleteData := db.Debug().Delete(&mappingRoleMenu)

	if deleteData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &mappingRoleMenu, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &mappingRoleMenu, <-err
}

/**
* ========================================================
* Repository Update Mapping Role Menu By ID Teritory
*=========================================================
 */

func (r *repositoryMappingRoleMenu) EntityUpdate(input *schemes.MappingRoleMenu) (*models.MappingRoleMenu, schemes.SchemeDatabaseError) {
	var mappingRoleMenu models.MappingRoleMenu
	mappingRoleMenu.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&mappingRoleMenu)

	checkId := db.Debug().First(&mappingRoleMenu)

	if checkId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &mappingRoleMenu, <-err
	}

	mappingRoleMenu.MerchantID = input.MerchantID
	mappingRoleMenu.RoleID = input.RoleID
	mappingRoleMenu.MenuID = input.MenuID
	mappingRoleMenu.MenuDetailID = input.MenuDetailID
	mappingRoleMenu.MenuDetailFunctionID = input.MenuDetailFunctionID
	mappingRoleMenu.Name = input.Name
	mappingRoleMenu.Active = input.Active

	updateData := db.Debug().Updates(&mappingRoleMenu)

	if updateData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &mappingRoleMenu, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &mappingRoleMenu, <-err
}
