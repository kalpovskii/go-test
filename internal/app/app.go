package app

import (
	"awesomeProject/internal/bootstrap"
	"awesomeProject/internal/config"
)

func Run(cfg config.Config) error {
	db, err := bootstrap.InitGormDB(cfg)
	if err != nil {
		return err
	}

	socksService := socksservice.New()
}
