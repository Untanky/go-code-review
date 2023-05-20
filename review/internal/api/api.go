package api

import (
	"context"
	"coupon_service/internal/config"
	"coupon_service/internal/entity"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Service interface {
	ApplyCoupon(entity.Basket, string) (*entity.Basket, error)
	CreateCoupon(int, string, int) any
	GetCoupons([]string) ([]entity.Coupon, error)
}

type api struct {
	srv *http.Server
	gin *gin.Engine
	svc Service
	cfg config.ApiConfig
}

func New(cfg config.ApiConfig, svc Service) *api {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	api := &api{
		gin: r,
		cfg: cfg,
		svc: svc,
	}
	return api.withServer().withRoutes()
}

func (a *api) withServer() *api {

	ch := make(chan *api)
	go func() {
		a.srv = &http.Server{
			Addr:    fmt.Sprintf(":%d", a.cfg.Port),
			Handler: a.gin,
		}
		ch <- a
	}()

	return <-ch
}

func (a *api) withRoutes() *api {
	apiGroup := a.gin.Group("/api")
	apiGroup.POST("/apply", a.ApplyCoupon)
	apiGroup.POST("/create", a.CreateCoupon)
	apiGroup.GET("/coupons", a.GetCoupon)
	return a
}

func (a *api) Start() {
	if err := a.srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func (a *api) Close() {
	<-time.After(5 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.srv.Shutdown(ctx); err != nil {
		log.Println(err)
	}
}
