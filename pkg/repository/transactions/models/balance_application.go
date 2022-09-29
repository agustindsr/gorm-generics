package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm-with-generics/pkg/models"
	"time"
)

type BalanceApplicationDAO struct {
	ID                        int64  `gorm:"primaryKey,column:id"`
	TransactionID             string `gorm:"column:transaction_id"`
	UserBalance               int64  `gorm:"column:user_balance"`
	PromotionalBalance        int64  `gorm:"column:promotional_balance"`
	WinningsBalance           int64  `gorm:"column:winnings_balance"`
	UserBalanceApplied        int64  `gorm:"column:user_balance_applied"`
	PromotionalBalanceApplied int64  `gorm:"column:promotional_balance_applied"`
	WinningsBalanceApplied    int64  `gorm:"column:winnings_balance_applied"`
}

func toBalanceApplicationDAO(m models.TransactionLedger) *BalanceApplicationDAO {
	return &BalanceApplicationDAO{
		ID:                        time.Time.UnixNano(time.Now()),
		TransactionID:             m.ID,
		UserBalance:               m.Metadata.UserBalance,
		PromotionalBalance:        m.Metadata.PromotionalBalance,
		WinningsBalance:           m.Metadata.WinningsBalance,
		UserBalanceApplied:        m.Metadata.UserBalanceApplied,
		PromotionalBalanceApplied: m.Metadata.PromotionalBalanceApplied,
		WinningsBalanceApplied:    m.Metadata.WinningsBalanceApplied,
	}
}

func toPackageModel(pa PackageApplicationDAO) *models.PaymentPackage {
	packageID, _ := primitive.ObjectIDFromHex(pa.PackageID)

	createdAt := primitive.NewDateTimeFromTime(pa.CreatedAt)

	return &models.PaymentPackage{
		ID:              &packageID,
		Name:            pa.Name,
		Price:           pa.Price,
		Credits:         pa.Credits,
		Sort:            pa.Sort,
		Currency:        pa.Currency,
		Image:           pa.Image,
		ImageWithCoupon: pa.ImageWithCoupon,
		Category:        pa.Category,
		Active:          pa.Active,
		CreatedAt:       &createdAt,
	}
}

func (ba BalanceApplicationDAO) TableName() string {
	return "balance_applications"
}
