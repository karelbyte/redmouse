package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var parameters interface{}

func ManagerFrontRoutes(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", parameters)
}
