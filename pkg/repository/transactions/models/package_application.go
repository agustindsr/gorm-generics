package models

import (
	"gorm-with-generics/pkg/models"
	"time"
)

type PackageApplicationDAO struct {
	ID              int64     `gorm:"primaryKey,column:id"`
	TransactionID   string    `gorm:"column:transaction_id"`
	PackageID       string    `gorm:"column:package_id"`
	Name            string    `gorm:"column:name"`
	Price           int64     `gorm:"column:price"`
	Credits         int64     `gorm:"column:credits"`
	Sort            int       `gorm:"column:sort"`
	Currency        string    `gorm:"column:currency"`
	Image           string    `gorm:"column:image"`
	ImageWithCoupon string    `gorm:"column:image_with_coupon"`
	Category        string    `gorm:"column:category"`
	Active          bool      `gorm:"column:active"`
	CreatedAt       time.Time `gorm:"column:created_at"`
}

func toPackageApplicationDAO(m models.TransactionLedger) *PackageApplicationDAO {
	if m.Metadata.Package == nil {
		return nil
	}

	p := *m.Metadata.Package

	return &PackageApplicationDAO{
		ID:              time.Time.UnixNano(time.Now()),
		TransactionID:   m.ID,
		PackageID:       p.ID.Hex(),
		Name:            p.Name,
		Price:           p.Price,
		Credits:         p.Credits,
		Sort:            p.Sort,
		Currency:        p.Currency,
		Image:           p.Image,
		ImageWithCoupon: p.ImageWithCoupon,
		Category:        p.Category,
		Active:          p.Active,
	}
}

func (pa PackageApplicationDAO) TableName() string {
	return "package_applications"
}
