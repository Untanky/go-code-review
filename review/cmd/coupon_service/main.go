package main

import (
	"coupon_service/internal/api"
	"coupon_service/internal/config"
	"coupon_service/internal/repository/memdb"
	"coupon_service/internal/service"
	"log"
	"time"
)

var (
	cfg  = config.New()
	repo = memdb.New()
)

const (
	day  = 24 * time.Hour
	year = 365 * day
)

func main() {
	svc := service.New(repo)
	api := api.New(cfg.API, svc)
	log.Println("Starting Coupon service server")
	api.Start()
	defer api.Close()
	log.Println("Started Coupon service server")
	<-time.After(1 * year)
	log.Println("Coupon service server alive for a year, closing")
}
