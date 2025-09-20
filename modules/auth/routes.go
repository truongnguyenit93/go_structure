package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func RegisterRoutes(router *gin.Engine, injector *do.Injector) {
	authRoute := router.Group("/api/v1/auth")
	{
		_ = authRoute
	}

}
