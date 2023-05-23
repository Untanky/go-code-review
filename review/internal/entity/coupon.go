package entity

type Coupon struct {
	ID             string
	Code           string `json:"code" binding:"required"`
	Discount       int    `json:"discount" binding:"required"`
	MinBasketValue int    `json:"minBasketValue" binding:"required"`
}
