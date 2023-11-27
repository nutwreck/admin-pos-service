package repositories

import (
	"net/http"
	"net/url"
	"sort"
	"strings"

	"github.com/nutwreck/admin-pos-service/configs"
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
* ====================================================
* Repository Results All Master Menu Relation Teritory
*=====================================================
 */

func (r *repositoryMenu) EntityGetMenu(input *schemes.Menu) (*[]schemes.GetMenu, schemes.SchemeDatabaseError) {
	var (
		menu            []models.Menu
		result          []schemes.GetMenu
		args            []interface{}
		sortData        string = "ASC"
		queryData       string = constants.EMPTY_VALUE
		queryAdditional string = constants.EMPTY_VALUE
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&menu)

	if input.Sort != constants.EMPTY_VALUE {
		unScape, _ := url.QueryUnescape(input.Sort)
		sortData = strings.Replace(unScape, "'", constants.EMPTY_VALUE, -1)
	}

	//Untuk mengambil detail data
	queryData = `
		SELECT
			menu.id,
			menu.name,
			menu.active
		FROM master.menus menu 
	`

	queryAdditional = `
		JOIN master.merchants merchant on menu.merchant_id = merchant.id
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

	queryAdditional += ` ORDER BY menu."name" ` + sortData

	getDatas := db.Raw(queryData+queryAdditional, args...).Scan(&result)

	if getDatas.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &result, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &result, <-err
}

func (r *repositoryMenu) EntityGetMenuDetail(input *schemes.MenuDetail) (*[]schemes.GetMenuDetail, schemes.SchemeDatabaseError) {
	var (
		menudetail      []models.MenuDetail
		result          []schemes.GetMenuDetail
		args            []interface{}
		sortData        string = "ASC"
		queryData       string = constants.EMPTY_VALUE
		queryAdditional string = constants.EMPTY_VALUE
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&menudetail)

	if input.Sort != constants.EMPTY_VALUE {
		unScape, _ := url.QueryUnescape(input.Sort)
		sortData = strings.Replace(unScape, "'", constants.EMPTY_VALUE, -1)
	}

	//Untuk mengambil detail data
	queryData = `
		SELECT
			menudetail.id,
			merchant.id as merchant_id,
			merchant.name as merchant_name,
			menu.id as menu_id,
			menu.name as menu_name,
			menudetail.name,
			menudetail.link,
			CASE
				WHEN menudetail.image != '' THEN CONCAT('` + configs.AccessFile + `',menudetail.image)
				ELSE ''
			END AS image,
			CASE
				WHEN menudetail.icon != '' THEN CONCAT('` + configs.AccessFile + `',menudetail.icon)
				ELSE ''
			END AS icon,
			menudetail.active
		FROM master.menu_details menudetail
	`

	queryAdditional = `
		JOIN master.menus menu on menudetail.menu_id = menu.id
		JOIN master.merchants merchant on menudetail.merchant_id = merchant.id
	`

	queryAdditional += ` WHERE TRUE`

	if input.MerchantID != constants.EMPTY_VALUE {
		queryAdditional += ` AND menudetail.merchant_id = ?`
		args = append(args, input.MerchantID)
	}

	if input.MenuID != constants.EMPTY_VALUE {
		queryAdditional += ` AND menudetail.menu_id = ?`
		args = append(args, input.MenuID)
	}

	queryAdditional += ` ORDER BY menudetail.name ` + sortData

	getDatas := db.Raw(queryData+queryAdditional, args...).Scan(&result)

	if getDatas.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &result, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &result, <-err
}

func (r *repositoryMenu) EntityGetMenuDetailFunction(input *schemes.MenuDetailFunction) (*[]schemes.GetMenuDetailFunction, schemes.SchemeDatabaseError) {
	var (
		menudetailfunction []models.MenuDetailFunction
		result             []schemes.GetMenuDetailFunction
		args               []interface{}
		sortData           string = "ASC"
		queryData          string = constants.EMPTY_VALUE
		queryAdditional    string = constants.EMPTY_VALUE
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&menudetailfunction)

	if input.Sort != constants.EMPTY_VALUE {
		unScape, _ := url.QueryUnescape(input.Sort)
		sortData = strings.Replace(unScape, "'", constants.EMPTY_VALUE, -1)
	}

	//Untuk mengambil detail data
	queryData = `
		SELECT
			merchant.id as merchant_id,
			merchant.name as merchant_name,
			menu.id as menu_id,
			menu.name as menu_name,
			menudetail.id as menu_detail_id,
			menudetail.name as menu_detail_name,
			menudetailfunction.id,
			menudetailfunction.name,
			menudetailfunction.link,
			menudetailfunction.active
		FROM master.menu_detail_functions menudetailfunction
	`

	queryAdditional = `
		JOIN master.menus menu on menudetailfunction.menu_id = menu.id
		JOIN master.menu_details menudetail on menudetailfunction.menu_detail_id = menudetail.id 
		JOIN master.merchants merchant on menudetailfunction.merchant_id = merchant.id
	`

	queryAdditional += ` WHERE TRUE`

	if input.MerchantID != constants.EMPTY_VALUE {
		queryAdditional += ` AND menudetailfunction.merchant_id = ?`
		args = append(args, input.MerchantID)
	}

	if input.MenuDetailID != constants.EMPTY_VALUE {
		queryAdditional += ` AND menudetailfunction.menu_detail_id = ?`
		args = append(args, input.MenuDetailID)
	}

	if input.MenuID != constants.EMPTY_VALUE {
		queryAdditional += ` AND menudetailfunction.menu_id = ?`
		args = append(args, input.MenuID)
	}

	queryAdditional += ` ORDER BY menudetailfunction.name ` + sortData

	getDatas := db.Raw(queryData+queryAdditional, args...).Scan(&result)

	if getDatas.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &result, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &result, <-err
}

func (r *repositoryMenu) EntityResultRelations(input *schemes.Menu) (*[]schemes.GetMenuRelation, schemes.SchemeDatabaseError) {
	var (
		menu            []models.Menu
		resultRaw       []schemes.GetMenuRelationRaw
		result          []schemes.GetMenuRelation
		args            []interface{}
		sortData        string = "ASC"
		queryData       string = constants.EMPTY_VALUE
		queryAdditional string = constants.EMPTY_VALUE
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&menu)

	if input.Sort != constants.EMPTY_VALUE {
		unScape, _ := url.QueryUnescape(input.Sort)
		sortData = strings.Replace(unScape, "'", constants.EMPTY_VALUE, -1)
	}

	//Untuk mengambil detail data
	queryData = `
		SELECT
			menu.id as menu_id,
			menu."name" as menu_name,
			menu.active as menu_active,
			menudetail.id as menu_detail_id,
			menudetail."name" as menu_detail_name,
			menudetail.link  as menu_detail_link,
			CASE
				WHEN menudetail.image != '' THEN CONCAT('` + configs.AccessFile + `',menudetail.image)
				ELSE ''
			END AS menu_detail_image,
			CASE
				WHEN menudetail.icon != '' THEN CONCAT('` + configs.AccessFile + `',menudetail.icon)
				ELSE ''
			END AS menu_detail_icon,
			menudetail.active as menu_detail_active,
			menudetailfunction.id as menu_detail_function_id,
			menudetailfunction."name" as menu_detail_function_name,
			menudetailfunction.link as menu_detail_function_link,
			menudetailfunction.active as menu_detail_function_active
		FROM master.menus menu 
	`

	queryAdditional = `
		LEFT JOIN master.menu_details menudetail on menu.id = menudetail.menu_id 
		LEFT JOIN master.menu_detail_functions menudetailfunction on menudetail.id = menudetailfunction.menu_detail_id 
		JOIN master.merchants merchant on menu.merchant_id = merchant.id
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

	//queryAdditional += ` ORDER BY ` + sortData

	getDatas := db.Raw(queryData+queryAdditional, args...).Scan(&resultRaw)

	//Mapping Data
	if len(resultRaw) > 0 {
		groupedData := make(map[schemes.GroupMenuKey]schemes.GetMenuRelation)
		for _, item := range resultRaw {
			key := schemes.GroupMenuKey{
				MenuID:   item.MenuID,
				MenuName: item.MenuName,
			}

			if _, exists := groupedData[key]; !exists {
				groupedData[key] = schemes.GetMenuRelation{
					ID:     item.MenuID,
					Name:   item.MenuName,
					Active: item.MenuActive,
				}
			}

			if item.MenuDetailID == constants.EMPTY_VALUE && item.MenuDetailName == constants.EMPTY_VALUE && item.MenuDetailLink == constants.EMPTY_VALUE && item.MenuDetailImage == constants.EMPTY_VALUE && item.MenuDetailIcon == constants.EMPTY_VALUE && item.MenuDetailActive == nil {
				// All fields are empty, set ListDetail to nil
				continue
			}

			groupedItem := schemes.ListMenuDetail{
				ID:                 item.MenuDetailID,
				Name:               item.MenuDetailName,
				Link:               item.MenuDetailLink,
				Image:              item.MenuDetailImage,
				Icon:               item.MenuDetailIcon,
				Active:             item.MenuDetailActive,
				ListDetailFunction: nil, // Initialize with an empty array
			}

			if item.MenuDetailFunctionID != constants.EMPTY_VALUE || item.MenuDetailFunctionName != constants.EMPTY_VALUE || item.MenuDetailFunctionLink != constants.EMPTY_VALUE || item.MenuDetailFunctionActive != nil {
				groupedItem.ListDetailFunction = append(groupedItem.ListDetailFunction, schemes.ListMenuDetailFunction{
					ID:     item.MenuDetailFunctionID,
					Name:   item.MenuDetailFunctionName,
					Link:   item.MenuDetailFunctionLink,
					Active: item.MenuDetailFunctionActive,
				})
			}

			temp := groupedData[key]
			temp.ListDetail = append(temp.ListDetail, groupedItem)
			// Assign the modified struct back to the map
			groupedData[key] = temp
		}

		// Convert the map to a slice of GetMenuRelation
		for _, value := range groupedData {
			result = append(result, value)
		}

		// Sort the result based on label_group (Name)
		if strings.ToUpper(sortData) == "ASC" {
			sort.Slice(result, func(i, j int) bool {
				return result[i].Name < result[j].Name
			})
		} else if strings.ToUpper(sortData) == "DESC" {
			sort.Slice(result, func(i, j int) bool {
				return result[i].Name > result[j].Name
			})
		}
	}

	if getDatas.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &result, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &result, <-err
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
