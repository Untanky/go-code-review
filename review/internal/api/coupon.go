package api

import (
	apiEntity "coupon_service/internal/api/entity"
	"coupon_service/internal/entity"
	"log"
	"strings"

	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *api) ApplyCoupon(c *gin.Context) {
	applyRequest := apiEntity.ApplyCouponRequest{}
	if err := c.ShouldBindJSON(&applyRequest); err != nil {
		log.Printf("error binding apply coupon request: %v\n", err)
		c.Status(http.StatusBadRequest)
		return
	}

	basket, err := a.svc.ApplyCoupon(applyRequest.Basket, applyRequest.Code)
	if err != nil {
		log.Printf("error applying coupon: %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, basket)
}

func (a *api) CreateCoupon(c *gin.Context) {
	coupon := entity.Coupon{}
	if err := c.ShouldBindJSON(&coupon); err != nil {
		log.Printf("error binding create coupon request: %v\n", err)
		c.Status(http.StatusBadRequest)
		return
	}

	err := a.svc.CreateCoupon(coupon.Discount, coupon.Code, coupon.MinBasketValue)
	if err != nil {
		log.Printf("error creating coupon: %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func (a *api) getCouponsBody(c *gin.Context) {
	couponRequest := apiEntity.CouponRequest{}
	if err := c.ShouldBindJSON(&couponRequest); err != nil {
		log.Printf("error binding get coupon request: %v\n", err)
		c.Status(http.StatusBadRequest)
		return
	}

	a.getCoupons(c, couponRequest)
}

func (a *api) getCouponsQuery(c *gin.Context) {
	couponRequest := apiEntity.CouponRequest{}
	if err := c.ShouldBindQuery(&couponRequest); err != nil {
		log.Printf("error binding get coupon request: %v\n", err)
		c.Status(http.StatusBadRequest)
		return
	}

	// When binding from query, the codes are a comma separated string
	// in the first element of the slice
	// we should define the comma as the delimiter and split the string
	couponRequest.Codes = strings.Split(couponRequest.Codes[0], ",")

	a.getCoupons(c, couponRequest)
}

func (a *api) getCoupons(c *gin.Context, couponRequest apiEntity.CouponRequest) {
	coupons, err := a.svc.GetCoupons(couponRequest.Codes)
	if err != nil {
		log.Printf("error getting coupons: %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, coupons)
}
