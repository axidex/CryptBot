package main

import (
	"context"
	"flag"
	"github.com/axidex/CryptBot/internal/api"
	"github.com/axidex/CryptBot/internal/app"
	"github.com/axidex/CryptBot/internal/config"
	"github.com/axidex/CryptBot/internal/telegram"
	"github.com/axidex/CryptBot/pkg/logger"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"github.com/oklog/run"
	"os"
	"strconv"
	"syscall"
)

func main() {
	var debug bool
	flag.BoolVar(&debug, "debug", false, "Enable debug mode")
	flag.Parse()

	appLogger := logger.NewZapLogger()

	err := godotenv.Load()
	if err != nil {
		appLogger.Fatal("Error loading .env file")
	}

	botKey := os.Getenv("BOT_KEY")

	desHost := os.Getenv("DES_HOST")
	desPort, _ := strconv.Atoi(os.Getenv("DES_PORT"))

	apiPort, _ := strconv.Atoi(os.Getenv("API_PORT"))

	apiConfig := api.Config{Port: apiPort}

	desCfg := config.DesConfig{DesHost: desHost, DesPort: desPort}

	appLogger.Infof("desCfg: %+v", desCfg)

	client := telegram.NewClient(botKey, debug, appLogger, desCfg)

	desClient := resty.New().SetBaseURL("http://" + desCfg.DesHost + ":" + strconv.Itoa(desCfg.DesPort))
	apiApp := api.CreateApp(apiConfig, appLogger, desClient)

	ctx := context.Background()

	runG := run.Group{}
	runG.Add(app.HandleApp(ctx, apiApp))
	runG.Add(app.HandleApp(ctx, client))
	runG.Add(run.SignalHandler(context.TODO(), syscall.SIGINT, syscall.SIGTERM))
	err = runG.Run()
	if err != nil {
		appLogger.Fatal(err)
	}
}
