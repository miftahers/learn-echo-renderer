package main

import (
	config "learn-echo-renderer/Config"
	"learn-echo-renderer/route"
	"log"
)

func main() {
	config.InitConfig()
	config.InitDB()

	routePayload := &route.Payload{
		DBGorm: config.DB,
		Config: config.Cfg,
	}

	e := route.InitRoute(routePayload)

	log.Fatal(e.Start(config.Cfg.APIPort))
}
