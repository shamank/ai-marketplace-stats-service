package main

import (
	"flag"
	"github.com/shamank/ai-marketplace-stats-service/internal/app"
	"github.com/shamank/ai-marketplace-stats-service/internal/config"
	"os"
	"os/signal"
	"syscall"
)

const defaultConfigPath = "./configs/prod.yaml"

func main() {

	var cfgPath string

	flag.StringVar(&cfgPath, "cfg", "", "path to config file")
	flag.Parse()

	if cfgPath == "" {
		if cfgPathEnv := os.Getenv("CONFIG_PATH"); cfgPathEnv != "" {
			cfgPath = cfgPathEnv
		} else {
			cfgPath = defaultConfigPath
		}
	}

	cfg, err := config.LoadConfig(cfgPath)
	if err != nil {
		panic(err)
	}

	application := app.NewApp(cfg)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		if err := application.Run(); err != nil {
			panic(err)
		}
	}()

	<-quit

	application.Stop()

}
