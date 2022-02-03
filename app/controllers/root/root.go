package root

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"product-api/lib/response"
)

func Index(context *gin.Context) {
	response.Json(context, http.StatusOK, nil)
}
