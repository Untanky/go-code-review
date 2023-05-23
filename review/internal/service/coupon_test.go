package service

import (
	"coupon_service/internal/entity"
	"coupon_service/internal/repository/memdb"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		repo Repository
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		{"initialize service", args{repo: nil}, Service{repo: nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_ApplyCoupon(t *testing.T) {
	memdb := memdb.New()
	coupon := entity.Coupon{
		Discount:       10,
		Code:           "Superdiscount",
		MinBasketValue: 10,
	}
	memdb.Save(&coupon)
	type fields struct {
		repo Repository
	}
	type args struct {
		basket entity.Basket
		code   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantB   *entity.Basket
		wantErr bool
	}{
		{"Apply 10%", fields{memdb}, args{entity.Basket{Value: 100}, "Superdiscount"}, &entity.Basket{Value: 100, AppliedCode: "Superdiscount", AppliedDiscount: 10, ApplicationSuccessful: true}, false},
		{"Apply without min basket value", fields{memdb}, args{entity.Basket{Value: 9}, "Superdiscount"}, nil, true},
		{"Apply with basket value 0", fields{memdb}, args{entity.Basket{Value: 0}, "Superdiscount"}, &entity.Basket{Value: 0}, false},
		{"Apply with negative basket value", fields{memdb}, args{entity.Basket{Value: -1}, "Superdiscount"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				repo: tt.fields.repo,
			}
			gotB, err := s.ApplyCoupon(tt.args.basket, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApplyCoupon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotB, tt.wantB) {
				t.Errorf("ApplyCoupon() gotB = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestService_CreateCoupon(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		discount       int
		code           string
		minBasketValue int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   any
	}{
		{"Create 10% coupon", fields{memdb.New()}, args{10, "Superdiscount", 55}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				repo: tt.fields.repo,
			}

			err := s.CreateCoupon(tt.args.discount, tt.args.code, tt.args.minBasketValue)
			if err != tt.want {
				t.Errorf("CreateCoupon() error = %v, want %v", err, tt.want)
			}
		})
	}
}

func TestService_GetCoupons(t *testing.T) {
	memdb := memdb.New()
	coupon1 := entity.Coupon{
		Discount:       10,
		Code:           "Superdiscount",
		MinBasketValue: 10,
	}
	coupon2 := entity.Coupon{
		Discount:       20,
		Code:           "Hyperdiscount",
		MinBasketValue: 10,
	}
	memdb.Save(&coupon1)
	memdb.Save(&coupon2)
	type fields struct {
		repo Repository
	}
	type args struct {
		codes []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entity.Coupon
		wantErr bool
	}{
		{"get empty coupons", fields{memdb}, args{[]string{}}, []entity.Coupon{}, false},
		{"get a single coupon", fields{memdb}, args{[]string{"Superdiscount"}}, []entity.Coupon{coupon1}, false},
		{"get a two coupons", fields{memdb}, args{[]string{"Superdiscount", "Hyperdiscount"}}, []entity.Coupon{coupon1, coupon2}, false},
		{"get non existing coupon", fields{memdb}, args{[]string{"Megacoupon"}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				repo: tt.fields.repo,
			}
			got, err := s.GetCoupons(tt.args.codes)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCoupons() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCoupons() got = %v, want %v", got, tt.want)
			}
		})
	}
}
