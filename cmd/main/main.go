package main

import (
	"bot/internal/telegram"
	"bot/pkg/logger"
	"flag"
	"github.com/joho/godotenv"
	"os"
	"os/signal"
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

	client := telegram.NewClient(botKey, debug, appLogger)
	go client.Run()

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	//registers the channel
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		sig := <-sigs
		appLogger.Warnf("Caught %s, shutting down", sig.String())
		client.Stop()
		// Finish any outstanding requests, then...
		done <- true
	}()

	<-done
}
