package memdb_test

import (
	"coupon_service/internal/entity"
	. "coupon_service/internal/repository/memdb"
	"reflect"
	"strings"
	"testing"
)

func TestRepository_FindByCode(t *testing.T) {
	repo := New()
	if repo == nil {
		t.Errorf("New() = %v, want %v", repo, nil)
		return
	}

	coupon := entity.Coupon{
		Discount:       10,
		Code:           "test",
		MinBasketValue: 100,
	}
	err := repo.Save(&coupon)
	if err != nil {
		t.Errorf("Save() error = %v, want %v", err, nil)
		return
	}

	foundCoupon, err := repo.FindByCode("test")
	if err != nil {
		t.Errorf("FindByCode() error = %v, want %v", err, nil)
		return
	}
	if !reflect.DeepEqual(foundCoupon, &coupon) {
		t.Errorf("FindByCode() foundCoupon = %v, want %v", *foundCoupon, coupon)
		return
	}
}

func TestRepository_FindByCode_UnknownCode(t *testing.T) {
	repo := New()
	if repo == nil {
		t.Errorf("New() = %v, want %v", repo, nil)
		return
	}

	foundCoupon, err := repo.FindByCode("test")
	if foundCoupon != nil {
		t.Errorf("FindByCode() foundCoupon = %v, want %v", *foundCoupon, nil)
		return
	}
	if !strings.Contains(err.Error(), "coupon not found") {
		t.Errorf("FindByCode() error = %v, want %v", err, nil)
		return
	}
}

func TestRepository_FindByCode_ShouldNotUpdateCouponWithoutSave(t *testing.T) {
	repo := New()

	coupon := entity.Coupon{
		Discount:       10,
		Code:           "test",
		MinBasketValue: 100,
	}
	repo.Save(&coupon)

	foundCoupon1, _ := repo.FindByCode("test")
	foundCoupon1.Discount = 20

	foundCoupon2, _ := repo.FindByCode("test")
	if foundCoupon2.Discount != 10 {
		t.Errorf("FindByCode() foundCoupon2 = %v, want %v", foundCoupon2, coupon)
		return
	}
}
