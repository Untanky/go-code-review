package entity

type Basket struct {
	Value                 int  `json:"value" binding:"required"`
	AppliedDiscount       int  `json:"applied_discount" binding:"required"`
	ApplicationSuccessful bool `json:"application_successful" binding:"required"`
}
