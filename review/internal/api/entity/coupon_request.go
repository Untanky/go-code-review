package entity

type CouponRequest struct {
	Codes []string `json:"codes" binding:"required" form:"codes"`
}
