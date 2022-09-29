package models

import "time"

const (
	PaymentPackageDefaultSort     = 1
	PaymentPackageDefaultCurrency = "MXN"

	PaymentPackageCategoryGeneral        = "general"
	PaymentPackageCategoryFirstTimePromo = "first_time_promo"

	PaymentPackageDefaultCategory = PaymentPackageCategoryGeneral
)

type PaymentPackage struct {
	ID              string     `json:"id"`
	Name            string     `json:"name"`
	Price           int64      `json:"price"`
	Credits         int64      `json:"credits"`
	Sort            int        `json:"sort"`
	Currency        string     `json:"currency"`
	Image           string     `json:"image"`
	ImageWithCoupon string     `json:"imageWithCoupon"`
	Category        string     `json:"category"`
	Active          bool       `json:"active"`
	CreatedAt       *time.Time `json:"createdAt"`
	UpdatedAt       *time.Time `json:"updatedAt"`
}
