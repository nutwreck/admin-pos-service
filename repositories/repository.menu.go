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

type repositoryMenu struct {
	db *gorm.DB
}

func NewRepositoryMenu(db *gorm.DB) *repositoryMenu {
	return &repositoryMenu{db: db}
}

/**
* ===============================================
* Repository Create New Master Menu Teritory
*================================================
 */

func (r *repositoryMenu) EntityCreate(input *schemes.Menu) (*models.Menu, schemes.SchemeDatabaseError) {
	var menu models.Menu
	menu.MerchantID = input.MerchantID
	menu.Name = input.Name

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&menu)

	checkMenuName := db.Debug().Where("merchant_id = ? AND name = ?", menu.MerchantID, menu.Name).First(&menu)

	if checkMenuName.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &menu, <-err
	}

	addmenu := db.Debug().Create(&menu).Commit()

	if addmenu.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_02",
		}
		return &menu, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &menu, <-err
}

/**
* ================================================
* Repository Results All Master Menu Teritory
*=================================================
 */

func (r *repositoryMenu) EntityResults(input *schemes.Menu) (*[]schemes.GetMenu, int64, schemes.SchemeDatabaseError) {
	var (
		menu            []models.Menu
		result          []schemes.GetMenu
		countData       schemes.CountData
		args            []interface{}
		totalData       int64
		sortData        string = "menu.created_at DESC"
		queryCountData  string = constants.EMPTY_VALUE
		queryData       string = constants.EMPTY_VALUE
		queryAdditional string = constants.EMPTY_VALUE
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&menu)

	if input.Sort != constants.EMPTY_VALUE {
		unScape, _ := url.QueryUnescape(input.Sort)
		sortData = strings.Replace(unScape, "'", constants.EMPTY_VALUE, -1)
	}

	offset := int((input.Page - 1) * input.PerPage)

	//Untuk mengambil jumlah data tanpa limit
	queryCountData = `
		SELECT
			COUNT(menu.*) AS count_data
		FROM master.menus AS menu
	`

	//Untuk mengambil detail data
	queryData = `
		SELECT
			menu.id,
			menu.name,
			menu.active,
			merchant.id AS merchant_id,
			merchant.name AS merchant_name,
			menu.created_at
		FROM master.menus AS menu
	`

	queryAdditional = `
		JOIN master.merchants AS merchant ON menu.merchant_id = merchant.id AND merchant.active = true
	`

	queryAdditional += ` WHERE TRUE`

	if input.MerchantID != constants.EMPTY_VALUE {
		queryAdditional += ` AND menu.merchant_id = ?`
		args = append(args, input.MerchantID)
	}

	if input.Name != constants.EMPTY_VALUE {
		queryAdditional += ` AND menu.name LIKE ?`
		args = append(args, "%"+input.Name+"%")
	}

	if input.ID != constants.EMPTY_VALUE {
		queryAdditional += ` AND menu.id = ?`
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
* =================================================
* Repository Delete Master Menu By ID Teritory
*==================================================
 */

func (r *repositoryMenu) EntityDelete(input *schemes.Menu) (*models.Menu, schemes.SchemeDatabaseError) {
	var menu models.Menu
	menu.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&menu)

	checkMenuId := db.Debug().First(&menu)

	if checkMenuId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &menu, <-err
	}

	deleteMenu := db.Debug().Delete(&menu)

	if deleteMenu.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &menu, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &menu, <-err
}

/**
* =================================================
* Repository Update Master Menu By ID Teritory
*==================================================
 */

func (r *repositoryMenu) EntityUpdate(input *schemes.Menu) (*models.Menu, schemes.SchemeDatabaseError) {
	var menu models.Menu
	menu.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&menu)

	checkMenuId := db.Debug().First(&menu)

	if checkMenuId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &menu, <-err
	}

	menu.MerchantID = input.MerchantID
	menu.Name = input.Name
	menu.Active = input.Active

	updateMenu := db.Debug().Updates(&menu)

	if updateMenu.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &menu, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &menu, <-err
}
