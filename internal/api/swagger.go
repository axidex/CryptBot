package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func (app *App) swagger(c *gin.Context) {
	switch c.Param("any") {
	case "/", "/docs":
		c.Redirect(http.StatusMovedPermanently, "/api/swagger/index.html")
	default:
		ginSwagger.WrapHandler(swaggerFiles.Handler)(c)
	}

}
