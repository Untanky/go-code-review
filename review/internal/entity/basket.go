package entity

type Basket struct {
	Value                 int    `json:"value" binding:"required"`
	AppliedCode           string `json:"applied_code"`
	AppliedDiscount       int    `json:"appliedDiscount"`
	ApplicationSuccessful bool   `json:"applicationSuccessful"`
}
