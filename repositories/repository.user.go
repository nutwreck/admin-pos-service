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
* Repository Register Auth Teritory
*===========================================
 */

func (r *repositoryUser) EntityRegister(input *schemes.User) (*models.User, schemes.SchemeDatabaseError) {
	var user models.User
	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password
	user.RoleID = input.RoleID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&user)

	checkEmailExist := db.Debug().First(&user, "email = ?", input.Email)

	if checkEmailExist.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_register_01",
		}
		return &user, <-err
	}

	addNewUser := db.Debug().Create(&user).Commit()

	if addNewUser.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_register_02",
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
