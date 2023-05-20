package entity

type CouponRequest struct {
	Codes []string `query:"codes"`
}
