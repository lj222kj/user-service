package main

import (
	"fmt"
	"user-service/database"
	"user-service/global"
	"user-service/users"
	"user-service/web"
)

func main() {
	cfg, err := global.New()
	if err != nil {
		fmt.Errorf("failed to read env config")
	}
	db := database.New()
	userSvc := users.New(db)
	api := web.New(cfg, userSvc)

	err = api.ListenAndServe(); if err != nil {
		fmt.Errorf("failed to listen to port", cfg.Port)
	}
}