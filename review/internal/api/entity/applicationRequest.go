package entity

import "coupon_service/internal/entity"

type ApplicationRequest struct {
	Code   string
	Basket entity.Basket
}
