package main

import (
	"flag"
	"log"
	"web-10/internal/hello/api"
	"web-10/internal/hello/config"
	"web-10/internal/hello/provider"
	"web-10/internal/hello/usecase"

	_ "github.com/lib/pq"
)

func main() {
	// Считываем аргументы командной строки
	configPath := flag.String("config-path", "configs/hello_example.yaml", "путь к файлу конфигурации")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	prv := provider.NewProvider(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBname)
	use := usecase.NewUsecase(cfg.Usecase.DefaultMessage, prv)
	srv := api.NewServer(cfg.IP, cfg.Port, cfg.API.MaxMessageSize, use)

	srv.Run()
}
