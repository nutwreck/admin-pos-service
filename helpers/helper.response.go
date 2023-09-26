package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/nutwreck/admin-pos-service/schemes"
)

func APIResponse(ctx *gin.Context, Message string, StatusCode int, Data interface{}) {

	jsonResponse := schemes.Responses{
		StatusCode: StatusCode,
		Message:    Message,
		Data:       Data,
	}

	if StatusCode >= 400 {
		ctx.AbortWithStatusJSON(StatusCode, jsonResponse)
	} else {
		ctx.JSON(StatusCode, jsonResponse)
	}
}

func APIResponsePagination(ctx *gin.Context, Message string, StatusCode int, Data interface{}, Page int, PerPage int, TotalPage int, TotalData int) {
	jsonResponse := schemes.ResponsesPagination{
		StatusCode: StatusCode,
		Message:    Message,
		Page:       Page,
		PerPage:    PerPage,
		TotalPage:  TotalPage,
		TotalData:  TotalData,
		Data:       Data,
	}

	if StatusCode >= 400 {
		ctx.AbortWithStatusJSON(StatusCode, jsonResponse)
	} else {
		ctx.JSON(StatusCode, jsonResponse)
	}
}

func ErrorResponse(ctx *gin.Context, Error interface{}) {
	var (
		data              schemes.ReadMsgErrorValidator
		errorsWithoutKeys []schemes.ResultMsgErrorValidator
	)

	if err := json.Unmarshal([]byte(Strigify(Error)), &data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": "500", "message": "Check Validator Data Error"})
		return
	}

	for _, err := range data.Results.Errors {
		for _, value := range err {
			errorsWithoutKeys = append(errorsWithoutKeys, value)
		}
	}

	err := schemes.ErrorResponse{
		StatusCode: http.StatusBadRequest,
		Error:      errorsWithoutKeys,
	}

	ctx.AbortWithStatusJSON(err.StatusCode, err)
}
