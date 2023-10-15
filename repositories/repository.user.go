package repositories

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/nutwreck/admin-pos-service/constants"
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/pkg"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type repositoryUser struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *repositoryUser {
	return &repositoryUser{db: db}
}

/**
* ==========================================
* Repository Add New User Auth Teritory
*===========================================
 */

func (r *repositoryUser) EntityAddUser(input *schemes.User) (*models.User, schemes.SchemeDatabaseError) {
	var user models.User
	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password
	user.MerchantID = input.MerchantID
	user.RoleID = input.RoleID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&user)

	checkEmailExist := db.Debug().First(&user, "email = ?", input.Email)

	if checkEmailExist.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_add_user_01",
		}
		return &user, <-err
	}

	addNewUser := db.Debug().Create(&user).Commit()

	if addNewUser.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_add_user_02",
		}
		return &user, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &user, <-err
}

/**
* ==========================================
* Repository Login Auth Teritory
*===========================================
 */

func (r *repositoryUser) EntityLogin(input *schemes.User) (*models.User, schemes.SchemeDatabaseError) {
	var user models.User
	user.Email = input.Email
	user.Password = input.Password

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&user)

	checkEmailExist := db.Debug().First(&user, "email = ?", input.Email)

	if checkEmailExist.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_login_01",
		}
		return &user, <-err
	}

	checkPasswordMatch := pkg.ComparePassword(user.Password, input.Password)

	if checkPasswordMatch != nil {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusBadRequest,
			Type: "error_login_02",
		}
		return &user, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &user, <-err
}

/**
* ==========================================
* Repository Get User Auth By ID Teritory
*===========================================
 */

func (r *repositoryUser) EntityGetUser(input *schemes.User) (*models.User, schemes.SchemeDatabaseError) {
	var user models.User
	user.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&user)

	checkExist := db.Debug().Where("id = ? AND active = ?", input.ID, true).First(&user)

	if checkExist.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &user, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &user, <-err
}

/**
* =================================================
* Repository Update User By ID Teritory
*==================================================
 */

func (r *repositoryUser) EntityUpdate(input *schemes.UpdateUser) (*models.User, schemes.SchemeDatabaseError) {
	var (
		user           models.User
		oldPassword    string
		newPassword    string
		changePassword string
	)

	user.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	//Pengecekan password input
	oldPassword = input.OldPassword
	newPassword = input.NewPassword
	if oldPassword != constants.EMPTY_VALUE && newPassword == constants.EMPTY_VALUE {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusBadRequest,
			Type: "error_update_03",
		}
		return &user, <-err
	}
	if oldPassword == constants.EMPTY_VALUE && newPassword != constants.EMPTY_VALUE {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusBadRequest,
			Type: "error_update_04",
		}
	}
	if oldPassword != constants.EMPTY_VALUE && newPassword != constants.EMPTY_VALUE {
		checkPasswordMatch := pkg.ComparePassword(input.DataPassword, oldPassword)
		if checkPasswordMatch != nil {
			err <- schemes.SchemeDatabaseError{
				Code: http.StatusBadRequest,
				Type: "error_update_05",
			}
			return &user, <-err
		}
		changePassword = newPassword
	}

	db := r.db.Model(&user)

	user.Name = input.Name
	user.RoleID = input.RoleID
	user.MerchantID = input.MerchantID
	user.Active = input.Active
	if changePassword != constants.EMPTY_VALUE {
		user.Password = changePassword
		updateUser := db.Debug().Updates(&user)

		if updateUser.RowsAffected < 1 {
			err <- schemes.SchemeDatabaseError{
				Code: http.StatusForbidden,
				Type: "error_update_02",
			}
			return &user, <-err
		}
	} else {
		updateUser := db.Debug().Model(&user).Omit("Password").Updates(&user)

		if updateUser.RowsAffected < 1 {
			err <- schemes.SchemeDatabaseError{
				Code: http.StatusForbidden,
				Type: "error_update_02",
			}
			return &user, <-err
		}
	}

	err <- schemes.SchemeDatabaseError{}
	return &user, <-err
}

/**
* =================================================
* Repository Result Master Role By ID Teritory
*==================================================
 */
func (r *repositoryUser) EntityGetRole(input *schemes.Role) (*models.Role, schemes.SchemeDatabaseError) {
	var role models.Role
	role.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&role)

	getData := db.Debug().Where("id = ? AND active = ?", input.ID, true).First(&role)

	if getData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &role, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &role, <-err
}

/**
* =================================================
* Repository Result Merchant By ID Teritory
*==================================================
 */
func (r *repositoryUser) EntityGetMerchant(input *schemes.Merchant) (*models.Merchant, schemes.SchemeDatabaseError) {
	var merchant models.Merchant
	merchant.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&merchant)

	checkMerchant := db.Debug().First(&merchant)

	if checkMerchant.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &merchant, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &merchant, <-err
}

/**
* =================================================
* Repository Result User Outlet Teritory
*==================================================
 */
func (r *repositoryUser) EntityGetUserOutlet(input *schemes.UserOutlet) (*[]schemes.GetUserOutlet, schemes.SchemeDatabaseError) {
	var (
		userOutlet      models.UserOutlet
		result          []schemes.GetUserOutlet
		args            []interface{}
		queryData       string = constants.EMPTY_VALUE
		queryAdditional string = constants.EMPTY_VALUE
	)
	userOutlet.UserID = input.UserID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&userOutlet)

	queryData = `
		SELECT
			user_outlet.id,
			user_outlet.user_id,
			users.name AS user_name,
			user_outlet.outlet_id,
			outlet.name AS outlet_name,
			outlet.phone AS outlet_phone,
			outlet.address AS outlet_address,
			outlet.description AS outlet_description,
			outlet.active AS outlet_active,
			outlet.created_at AS outlet_created_at
		FROM master.user_outlets AS user_outlet
	`

	queryAdditional = `
		JOIN master.users ON user_outlet.user_id = users.id AND users.active = true
		JOIN master.outlets AS outlet ON user_outlet.outlet_id = outlet.id AND outlet.active = true
	`

	queryAdditional += ` WHERE TRUE`

	if input.UserID != constants.EMPTY_VALUE {
		queryAdditional += ` AND user_outlet.user_id = ?`
		args = append(args, input.UserID)
	}

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
