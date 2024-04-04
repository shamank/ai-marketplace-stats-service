package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/shamank/ai-marketplace-stats-service/internal/app"
	"github.com/shamank/ai-marketplace-stats-service/internal/config"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	var cfgPath string

	flag.StringVar(&cfgPath, "cfg", "", "path to config file")
	flag.Parse()

	if err := godotenv.Load(); err != nil {
		fmt.Println("error occured while loading .env file, error: ", err) // нам не обязательно падать с ошибкой
	}

	cfg, err := config.LoadConfig(cfgPath)
	if err != nil {
		panic(err)
	}

	application := app.NewApp(cfg)

	quit := make(chan os.Signal, 1) // канал для реализации graceful shutdown
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		if err := application.Run(); err != nil {
			panic(err)
		}
	}()

	<-quit

	application.Stop()

}
