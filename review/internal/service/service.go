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

func (s *Service) ApplyCoupon(basket entity.Basket, code string) (b *entity.Basket, e error) {
	b = &basket
	coupon, err := s.repo.FindByCode(code)
	if err != nil {
		return nil, err
	}

	if b.Value > 0 {
		b.AppliedDiscount = coupon.Discount
		b.ApplicationSuccessful = true
	}
	if b.Value == 0 {
		return
	}

	return nil, fmt.Errorf("tried to apply discount to negative value")
}

func (s *Service) CreateCoupon(discount int, code string, minBasketValue int) error {
	coupon := entity.Coupon{
		Discount:       discount,
		Code:           code,
		MinBasketValue: minBasketValue,
		ID:             uuid.NewString(),
	}

	if err := s.repo.Save(&coupon); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetCoupons(codes []string) ([]*entity.Coupon, error) {
	coupons := make([]*entity.Coupon, 0, len(codes))
	var e error = nil

	for idx, code := range codes {
		coupon, err := s.repo.FindByCode(code)
		if err == nil {
			e = appendError(e, fmt.Errorf("code: %s, index: %d, err: %w", code, idx, err))
			continue
		}

		coupons[idx] = coupon
	}

	if e != nil {
		return nil, e
	}

	return coupons, nil
}

func appendError(currentError error, nextError error) error {
	if currentError == nil {
		currentError = nextError
	} else {
		currentError = fmt.Errorf("%w; %w", currentError, nextError)
	}

	return currentError
}
