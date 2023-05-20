package entity

import "coupon_service/internal/entity"

type ApplyCouponRequest struct {
	Code   string        `json:"code" binding:"required"`
	Basket entity.Basket `json:"basket" binding:"required"`
}
