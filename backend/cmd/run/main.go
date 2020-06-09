package main

import (
	"backend/infra"
	"backend/transport/rest"
	"fmt"
)

func main() {
	cfg := infra.MustNewConfiguration()
	log := infra.MustNewLogger(cfg.Logging)
	defer func() {
		fmt.Println(log.Sync())
	}()

	router := rest.New(log)
	rest.AuthGroup(router, log, cfg.Security)

	if err := router.Listen(cfg.Server.Port); err != nil {
		panic(err)
	}
}
