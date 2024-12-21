package api

import (
	"fmt"
	_ "github.com/axidex/CryptBot/docs"
	"github.com/axidex/CryptBot/internal/app"
	"github.com/axidex/CryptBot/pkg/logger"
	"github.com/gin-gonic/gin"
)

type App struct {
	config Config
	logger logger.Logger
	engine *gin.Engine
}

func CreateApp(config Config, logger logger.Logger) app.App {
	apiApp := &App{
		config: config,
		logger: logger,
	}
	apiApp.InitRoutes()
	return apiApp
}

func (app *App) InitRoutes() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(CustomRecoveryFunc(app.logger))
	router.Use(LoggerMiddleware(app.logger))

	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/swagger/*any", app.swagger)

	api := router.Group("/api")
	{
		health := api.Group("/health")
		{
			health.GET("/ping", app.health)
		}

		v1 := api.Group("/v1")
		{
			v1.Use(rateLimiter)
			v1.POST("/des3", app.Des3)
		}

	}

	app.listRoutes(router)
	app.engine = router
}

func (app *App) listRoutes(router *gin.Engine) {
	for _, route := range router.Routes() {
		app.logger.Infof("Method: %s | Path: %s | Handler: %s", route.Method, route.Path, route.Handler)
	}
}

func (app *App) Run() {
	for {
		err := app.engine.Run(fmt.Sprintf(":%d", app.config.Port))
		if err != nil {
			app.logger.Fatalf("Failed to start server - %s", err)
		}
	}
}

func (app *App) Stop() {

}
