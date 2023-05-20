package main

import (
	"coupon_service/internal/api"
	"coupon_service/internal/config"
	"coupon_service/internal/repository/memdb"
	"coupon_service/internal/service"
	"fmt"
	"time"
)

var (
	cfg  = config.New()
	repo = memdb.New()
)

const (
	day = time.Hour * 24
	year = day * 365
)

func main() {
	svc := service.New(repo)
	api := api.New(cfg.API, svc)
	api.Start()
	defer api.Close()
	fmt.Println("Starting Coupon service server")
	<-time.After(1 * year)
	fmt.Println("Coupon service server alive for a year, closing")
}
