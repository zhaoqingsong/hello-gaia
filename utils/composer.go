package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/log4go"
	"net/http"
)


// ListData .
type ListData struct {
	Page     int64       `json:"page"`
	Pagesize int64       `json:"pagesize"`
	Total    int64       `json:"total"`
	Items    interface{} `json:"items"`
}

// Pagination .
type Pagination struct {
	Page     int64 `json:"page"`
	Pagesize int64 `json:"pagesize"`
	Total    int64 `json:"total"`
}

// Resp .
type Resp struct {
	Result  int         `json:"result"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ComposerJSON .
func ComposerJSON(method string, ctx *gin.Context, err error, data interface{}, pagination *Pagination) {

	code := 1
	message := "success."

	if pagination != nil {

		if pagination.Page == 0 {
			pagination.Page = 1
		}
		if pagination.Pagesize == 0 {
			pagination.Pagesize = 20
		}

		data = &ListData{
			Page:     pagination.Page,
			Pagesize: pagination.Pagesize,
			Total:    pagination.Total,
			Items:    data,
		}
	}

	if err != nil {
		code = -1
		message = fmt.Sprint(err)
		log4go.Info("%s failed, err: %s", method, err)
	} else {
		log4go.Info("%s succeed, data: %v", method, data)
	}

	ctx.JSON(http.StatusOK, &Resp{
		Result:  code,
		Message: message,
		Data:    data,
	})

	return
}
