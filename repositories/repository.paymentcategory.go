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

type repositoryPaymentCategory struct {
	db *gorm.DB
}

func NewRepositoryPaymentCategory(db *gorm.DB) *repositoryPaymentCategory {
	return &repositoryPaymentCategory{db: db}
}

/**
* ======================================================
* Repository Create New Master Payment Category Teritory
*=======================================================
 */

func (r *repositoryPaymentCategory) EntityCreate(input *schemes.PaymentCategory) (*models.PaymentCategory, schemes.SchemeDatabaseError) {
	var paymentCategory models.PaymentCategory
	paymentCategory.MerchantID = input.MerchantID
	paymentCategory.Name = input.Name

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&paymentCategory)

	checkData := db.Debug().Where("merchant_id = ? AND name = ?", paymentCategory.MerchantID, paymentCategory.Name).First(&paymentCategory)

	if checkData.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &paymentCategory, <-err
	}

	addData := db.Debug().Create(&paymentCategory).Commit()

	if addData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_02",
		}
		return &paymentCategory, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &paymentCategory, <-err
}

/**
* =======================================================
* Repository Results All Master Payment Category Teritory
*========================================================
 */

func (r *repositoryPaymentCategory) EntityResults(input *schemes.PaymentCategory) (*[]schemes.GetPaymentCategory, int64, schemes.SchemeDatabaseError) {
	var (
		paymentCategory []models.PaymentCategory
		result          []schemes.GetPaymentCategory
		countData       schemes.CountData
		args            []interface{}
		totalData       int64
		sortData        string = "paymentcategory.created_at DESC"
		queryCountData  string = constants.EMPTY_VALUE
		queryData       string = constants.EMPTY_VALUE
		queryAdditional string = constants.EMPTY_VALUE
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&paymentCategory)

	if input.Sort != constants.EMPTY_VALUE {
		unScape, _ := url.QueryUnescape(input.Sort)
		sortData = strings.Replace(unScape, "'", constants.EMPTY_VALUE, -1)
	}

	offset := int((input.Page - 1) * input.PerPage)

	//Untuk mengambil jumlah data tanpa limit
	queryCountData = `
		SELECT
			COUNT(paymentcategory.*) AS count_data
		FROM master.payment_categorys AS paymentcategory
	`

	//Untuk mengambil detail data
	queryData = `
		SELECT
			paymentcategory.id,
			paymentcategory.name,
			paymentcategory.active,
			paymentcategory.created_at,
			merchant.id AS merchant_id,
			merchant.name AS merchant_name
		FROM master.payment_categorys AS paymentcategory
	`

	queryAdditional = `
		JOIN master.merchants AS merchant ON paymentcategory.merchant_id = merchant.id AND merchant.active = true
	`

	queryAdditional += ` WHERE TRUE`

	if input.MerchantID != constants.EMPTY_VALUE {
		queryAdditional += ` AND paymentcategory.merchant_id = ?`
		args = append(args, input.MerchantID)
	}

	if input.Name != constants.EMPTY_VALUE {
		queryAdditional += ` AND paymentcategory.name LIKE ?`
		args = append(args, "%"+input.Name+"%")
	}

	if input.ID != constants.EMPTY_VALUE {
		queryAdditional += ` AND paymentcategory.id = ?`
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
* Repository Delete Master Payment Category By ID Teritory
*=========================================================
 */

func (r *repositoryPaymentCategory) EntityDelete(input *schemes.PaymentCategory) (*models.PaymentCategory, schemes.SchemeDatabaseError) {
	var paymentCategory models.PaymentCategory
	paymentCategory.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&paymentCategory)

	checkId := db.Debug().First(&paymentCategory)

	if checkId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &paymentCategory, <-err
	}

	deleteData := db.Debug().Delete(&paymentCategory)

	if deleteData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &paymentCategory, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &paymentCategory, <-err
}

/**
* ========================================================
* Repository Update Master Payment Category By ID Teritory
*=========================================================
 */

func (r *repositoryPaymentCategory) EntityUpdate(input *schemes.PaymentCategory) (*models.PaymentCategory, schemes.SchemeDatabaseError) {
	var paymentCategory models.PaymentCategory
	paymentCategory.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&paymentCategory)

	checkId := db.Debug().First(&paymentCategory)

	if checkId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &paymentCategory, <-err
	}

	paymentCategory.MerchantID = input.MerchantID
	paymentCategory.Name = input.Name
	paymentCategory.Active = input.Active

	updateData := db.Debug().Updates(&paymentCategory)

	if updateData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &paymentCategory, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &paymentCategory, <-err
}
