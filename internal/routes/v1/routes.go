package v1

import (
	"github.com/gin-gonic/gin"
)

func RoutesV1(api *gin.RouterGroup) {
	basePath := "/v1"
	v1Group := api.Group(basePath)
	UserRoutesV1(v1Group)
}
