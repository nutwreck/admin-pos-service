package repositories

import (
	"net/http"
	"net/url"
	"strings"

	"gorm.io/gorm"

	"github.com/nutwreck/admin-pos-service/configs"
	"github.com/nutwreck/admin-pos-service/constants"
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type repositoryPaymentMethod struct {
	db *gorm.DB
}

func NewRepositoryPaymentMethod(db *gorm.DB) *repositoryPaymentMethod {
	return &repositoryPaymentMethod{db: db}
}

/**
* ====================================================
* Repository Create New Master Payment Method Teritory
*=====================================================
 */

func (r *repositoryPaymentMethod) EntityCreate(input *schemes.PaymentMethod) (*models.PaymentMethod, schemes.SchemeDatabaseError) {
	var paymentMethod models.PaymentMethod
	paymentMethod.Name = input.Name
	paymentMethod.MerchantID = input.MerchantID
	paymentMethod.PaymentCategoryID = input.PaymentCategoryID
	paymentMethod.AccountNumber = input.AccountNumber
	paymentMethod.Logo = input.Logo

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&paymentMethod)

	checkData := db.Debug().Where("merchant_id = ? AND payment_category_id = ? AND name = ?", paymentMethod.MerchantID, paymentMethod.PaymentCategoryID, paymentMethod.Name).First(&paymentMethod)

	if checkData.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &paymentMethod, <-err
	}

	addData := db.Debug().Create(&paymentMethod).Commit()

	if addData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_02",
		}
		return &paymentMethod, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &paymentMethod, <-err
}

/**
* =====================================================
* Repository Results All Master Payment Method Teritory
*======================================================
 */

func (r *repositoryPaymentMethod) EntityResults(input *schemes.PaymentMethod) (*[]schemes.GetPaymentMethod, int64, schemes.SchemeDatabaseError) {
	var (
		paymentMethod   []models.PaymentMethod
		result          []schemes.GetPaymentMethod
		countData       schemes.CountData
		args            []interface{}
		totalData       int64
		sortData        string = "paymentmethod.created_at DESC"
		queryCountData  string = constants.EMPTY_VALUE
		queryData       string = constants.EMPTY_VALUE
		queryAdditional string = constants.EMPTY_VALUE
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&paymentMethod)

	if input.Sort != constants.EMPTY_VALUE {
		unScape, _ := url.QueryUnescape(input.Sort)
		sortData = strings.Replace(unScape, "'", constants.EMPTY_VALUE, -1)
	}

	offset := int((input.Page - 1) * input.PerPage)

	//Untuk mengambil jumlah data tanpa limit
	queryCountData = `
		SELECT
			COUNT(paymentmethod.*) AS count_data
		FROM master.payment_methods AS paymentmethod
	`

	//Untuk mengambil detail data
	queryData = `
		SELECT
			paymentmethod.id,
			paymentmethod.name,
			paymentcategory.id AS payment_category_id,
			paymentcategory.name AS payment_category_name,
			COALESCE(paymentmethod.account_number,'') AS account_number,
			CASE
				WHEN paymentmethod.logo != '' THEN CONCAT('` + configs.AccessFile + `',paymentmethod.logo)
				ELSE ''
			END AS logo,
			paymentmethod.active,
			paymentmethod.created_at,
			merchant.id AS merchant_id,
			merchant.name AS merchant_name
		FROM master.payment_methods AS paymentmethod
	`

	queryAdditional = `
		JOIN master.merchants AS merchant ON paymentmethod.merchant_id = merchant.id AND merchant.active = true
		JOIN master.payment_categorys AS paymentcategory ON paymentmethod.payment_category_id = paymentcategory.id AND paymentcategory.active = true
	`

	queryAdditional += ` WHERE TRUE`

	if input.MerchantID != constants.EMPTY_VALUE {
		queryAdditional += ` AND paymentmethod.merchant_id = ?`
		args = append(args, input.MerchantID)
	}

	if input.PaymentCategoryID != constants.EMPTY_VALUE {
		queryAdditional += ` AND paymentmethod.payment_category_id = ?`
		args = append(args, input.PaymentCategoryID)
	}

	if input.Name != constants.EMPTY_VALUE {
		queryAdditional += ` AND paymentmethod.name LIKE ?`
		args = append(args, "%"+input.Name+"%")
	}

	if input.ID != constants.EMPTY_VALUE {
		queryAdditional += ` AND paymentmethod.id = ?`
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
* ======================================================
* Repository Result Master Payment Method By ID Teritory
*=======================================================
 */

func (r *repositoryPaymentMethod) EntityResult(input *schemes.PaymentMethod) (*models.PaymentMethod, schemes.SchemeDatabaseError) {
	var paymentMethod models.PaymentMethod
	paymentMethod.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&paymentMethod)

	checkData := db.Debug().First(&paymentMethod)

	if checkData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &paymentMethod, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &paymentMethod, <-err
}

/**
* ======================================================
* Repository Delete Master Payment Method By ID Teritory
*=======================================================
 */

func (r *repositoryPaymentMethod) EntityDelete(input *schemes.PaymentMethod) (*models.PaymentMethod, schemes.SchemeDatabaseError) {
	var paymentMethod models.PaymentMethod
	paymentMethod.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&paymentMethod)

	checkData := db.Debug().First(&paymentMethod)

	if checkData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &paymentMethod, <-err
	}

	deleteData := db.Debug().Delete(&paymentMethod)

	if deleteData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &paymentMethod, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &paymentMethod, <-err
}

/**
* ======================================================
* Repository Update Master Payment Method By ID Teritory
*=======================================================
 */

func (r *repositoryPaymentMethod) EntityUpdate(input *schemes.PaymentMethod) (*models.PaymentMethod, schemes.SchemeDatabaseError) {
	var paymentMethod models.PaymentMethod
	paymentMethod.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&paymentMethod)

	checkData := db.Debug().First(&paymentMethod)

	if checkData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &paymentMethod, <-err
	}

	paymentMethod.Name = input.Name
	paymentMethod.MerchantID = input.MerchantID
	paymentMethod.PaymentCategoryID = input.PaymentCategoryID
	paymentMethod.AccountNumber = input.AccountNumber
	paymentMethod.Logo = input.Logo
	paymentMethod.Active = input.Active

	updateData := db.Debug().Updates(&paymentMethod)

	if updateData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &paymentMethod, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &paymentMethod, <-err
}
