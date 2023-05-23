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
	router := gin.New()
	router.Use(gin.Recovery(), gin.Logger(), gin.ErrorLogger())

	api := &api{
		gin: router,
		cfg: cfg,
		svc: svc,
	}
	return api.withServer().withRoutesV1().withRoutesV2()
}

func (a *api) withServer() *api {
	a.srv = &http.Server{
		Addr:    fmt.Sprintf(":%d", a.cfg.Port),
		Handler: a.gin,
	}

	return a
}

func (a *api) withRoutesV1() *api {
	// we'll keep this routes for backward compatibility
	// but this does not follow RESTful API design
	apiGroup := a.gin.Group("/api")
	apiGroup.POST("/apply", a.ApplyCoupon)
	apiGroup.POST("/create", a.CreateCoupon)
	apiGroup.GET("/coupons", a.getCouponsBody)

	return a
}

func (a *api) withRoutesV2() *api {
	basicAuth := gin.BasicAuth(gin.Accounts{"admin": "super"})

	couponV2 := a.gin.Group("/api/v2/coupon")
	couponV2.POST("/apply", a.ApplyCoupon)
	couponV2.POST("", basicAuth, a.CreateCoupon)
	couponV2.GET("", basicAuth, a.getCouponsQuery) // depending on the use case, this does not need to be authenticated

	return a
}

func (a *api) Start() chan any {
	go func() {
		if err := a.srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	return make(chan any)
}

func (a *api) Close() {
	<-time.After(5 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.srv.Shutdown(ctx); err != nil {
		log.Println(err)
	}
}
