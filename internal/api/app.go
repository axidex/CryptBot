package api

import (
	"fmt"
	_ "github.com/axidex/CryptBot/docs"
	"github.com/axidex/CryptBot/internal/app"
	"github.com/axidex/CryptBot/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type App struct {
	config    Config
	logger    logger.Logger
	engine    *gin.Engine
	desClient *resty.Client
}

func CreateApp(config Config, logger logger.Logger, desClient *resty.Client) app.App {
	apiApp := &App{
		config:    config,
		logger:    logger,
		desClient: desClient,
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
			v1.POST("/a5", app.A5)
			v1.POST("/aes", app.Aes)
			v1.POST("/blowfish", app.Blowfish)
			v1.POST("/des", app.Des)
			v1.POST("/diffie", app.Diffie)
			v1.POST("/elgamal", app.Elgamal)
			v1.POST("/enigma", app.Enigma)
			v1.POST("/feistel", app.Feistel)
			v1.POST("/hash", app.Hash)
			v1.POST("/invMix", app.InvMix)
			v1.POST("/md5", app.Md5)
			v1.POST("/rsa", app.Rsa)
			v1.POST("/sBlock", app.SBlock)
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

func (app *App) Run() error {
	route := fmt.Sprintf(":%d", app.config.Port)
	app.logger.Infof("Starting api %s", route)
	return app.engine.Run(route)
}

func (app *App) Stop(err error) {
	app.logger.Infof("Stopping api")
	app.logger.Infof("Api stopped")
}
