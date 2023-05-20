package api

import (
	apiEntity "coupon_service/internal/api/entity"
	"coupon_service/internal/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *API) Apply(c *gin.Context) {
	apiReq := apiEntity.ApplicationRequest{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		return
	}
	basket, err := a.svc.ApplyCoupon(apiReq.Basket, apiReq.Code)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, basket)
}

func (a *API) Create(c *gin.Context) {
	apiReq := entity.Coupon{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		return
	}
	err := a.svc.CreateCoupon(apiReq.Discount, apiReq.Code, apiReq.MinBasketValue)
	if err != nil {
		return
	}
	c.Status(http.StatusOK)
}

func (a *API) Get(c *gin.Context) {
	apiReq := apiEntity.CouponRequest{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		return
	}
	coupons, err := a.svc.GetCoupons(apiReq.Codes)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, coupons)
}
