package main

import (
	"coupon_service/internal/api"
	"coupon_service/internal/config"
	"coupon_service/internal/repository/memdb"
	"coupon_service/internal/service"
	"log"
)

var (
	cfg  = config.New()
	repo = memdb.New()
)

func main() {
	svc := service.New(repo)
	api := api.New(cfg.API, svc)
	log.Println("Starting Coupon service server")
	ch := api.Start()
	defer api.Close()
	log.Println("Started Coupon service server")
	<-ch
}
