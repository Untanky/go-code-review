package service

import (
	"coupon_service/internal/entity"
	"fmt"

	"github.com/google/uuid"
)

type Repository interface {
	FindByCode(string) (*entity.Coupon, error)
	Save(*entity.Coupon) error
}

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) ApplyCoupon(basket entity.Basket, code string) (*entity.Basket, error) {
	b := &basket
	if b.Value < 0 {
		return nil, fmt.Errorf("tried to apply discount to negative value")
	}
	if b.Value == 0 {
		return b, nil
	}

	coupon, err := s.repo.FindByCode(code)
	if err != nil {
		return nil, err
	}

	if coupon.MinBasketValue > b.Value {
		return nil, fmt.Errorf("basket value is too low")
	}

	// FOLLOWUP: Do we need to subtract the discount from the basket value?
	b.AppliedCode = coupon.Code
	b.AppliedDiscount = coupon.Discount
	b.ApplicationSuccessful = true
	return b, nil
}

func (s *Service) CreateCoupon(discount int, code string, minBasketValue int) any {
	coupon := entity.Coupon{
		Discount:       discount,
		Code:           code,
		MinBasketValue: minBasketValue,
		ID:             uuid.NewString(),
	}

	return s.repo.Save(&coupon)
}

func (s *Service) GetCoupons(codes []string) ([]entity.Coupon, error) {
	coupons := make([]entity.Coupon, 0, len(codes))
	var e error = nil

	for idx, code := range codes {
		coupon, err := s.repo.FindByCode(code)
		if err != nil {
			e = appendError(e, CouponError{code, idx, err})
		} else {
			coupons = append(coupons, *coupon)
		}
	}
	if e != nil {
		return nil, e
	}

	return coupons, nil
}
