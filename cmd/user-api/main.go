package main

import (
	"fmt"
	"user-service/global"
	"user-service/web"
)

func main() {
	cfg, err := global.New()
	if err != nil {
		fmt.Errorf("failed to read env config")
	}
	api := web.New(cfg)

	err = api.ListenAndServe(); if err != nil {
		fmt.Errorf("failed to listen to port", cfg.Port)
	}
}
