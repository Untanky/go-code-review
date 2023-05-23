package entity

type Basket struct {
	Value                 int    `json:"value" binding:"required"`
	AppliedCode           string `json:"applied_code" binding:"required"`
	AppliedDiscount       int    `json:"applied_discount" binding:"required"`
	ApplicationSuccessful bool   `json:"application_successful" binding:"required"`
}
