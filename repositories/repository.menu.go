package repositories

import (
	"net/http"
	"net/url"
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
	menu.Name = input.Name

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&menu)

	checkMenuName := db.Debug().First(&menu, "name = ?", menu.Name)

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

func (r *repositoryMenu) EntityResults(input *schemes.Menu) (*[]models.Menu, int64, schemes.SchemeDatabaseError) {
	var (
		menu      []models.Menu
		totalData int64
		sort      string = configs.SortByDefault + " " + configs.OrderByDefault
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&menu)

	if input.Sort != constants.EMPTY_VALUE {
		unScape, _ := url.QueryUnescape(input.Sort)
		sort = strings.Replace(unScape, "'", constants.EMPTY_VALUE, -1)
	}

	if input.Name != constants.EMPTY_VALUE {
		db = db.Where("name LIKE ?", "%"+input.Name+"%")
	}

	if input.ID != constants.EMPTY_VALUE {
		db = db.Where("id LIKE ?", "%"+input.ID+"%")
	}

	offset := int((input.Page - 1) * input.PerPage)

	checkMenu := db.Debug().Order(sort).Offset(offset).Limit(int(input.PerPage)).Find(&menu)

	if checkMenu.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &menu, totalData, <-err
	}

	// Menghitung total data yang diambil
	db.Model(&models.Menu{}).Count(&totalData)

	err <- schemes.SchemeDatabaseError{}
	return &menu, totalData, <-err
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
