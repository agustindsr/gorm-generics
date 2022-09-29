package transactions

import (
	"gorm-with-generics/pkg/models"
	"time"
)

type transactionDAO struct {
	Id                        string            `gorm:"primaryKey"`
	WalletID                  string            `gorm:"column:wallet_id"`
	UserID                    string            `gorm:"column:user_id"`
	Amount                    int64             `gorm:"column:amount"`
	Currency                  string            `gorm:"column:currency"`
	Type                      string            `gorm:"column:type"`
	Kind                      string            `gorm:"column:kind"`
	IsReverse                 bool              `gorm:"column:is_reverse"`
	Status                    string            `gorm:"column:status"`
	Provider                  string            `gorm:"column:provider"`
	CreatedAt                 time.Time         `gorm:"column:created_at"`
	UpdatedAt                 time.Time         `gorm:"column:updated_at"`
	Token                     string            `gorm:"column:metadata_token"`
	DeviceID                  string            `gorm:"column:metadata_device_id"`
	UserCouponID              string            `gorm:"column:metadata_user_coupon_id"`
	OriginalPrice             int64             `gorm:"column:metadata_original_price"`
	OriginalCredits           int64             `gorm:"column:metadata_original_credits"`
	LineupID                  string            `gorm:"column:metadata_lineup_id"`
	ContestID                 string            `gorm:"column:metadata_contest_id"`
	UserBalance               int64             `gorm:"column:metadata_user_balance"`
	PromotionalBalance        int64             `gorm:"column:metadata_promotional_balance"`
	WinningsBalance           int64             `gorm:"column:metadata_winnings_balance"`
	UserBalanceApplied        int64             `gorm:"column:metadata_user_Balance_applied"`
	PromotionalBalanceApplied int64             `gorm:"column:metadata_promotional_balance_applied"`
	WinningsBalanceApplied    int64             `gorm:"column:metadata_winnings_Balance_applied"`
	Rake                      int64             `gorm:"column:metadata_rake"`
	RakePercentage            float64           `gorm:"column:metadata_rake_percentage"`
	OriginalReferenceID       string            `gorm:"column:metadata_original_reference_id"`
	ReverseReason             string            `gorm:"column:metadata_reverse_reason"`
	Clabe                     string            `gorm:"column:metadata_clabe"`
	BankAccountInformationID  string            `gorm:"column:metadata_bank_account_information_id"`
	PaymentPackageDAO         PaymentPackageDAO `gorm:"foreignKey:TransactionID"`
}

type PaymentPackageDAO struct {
	ID               int    `gorm:"column:id"`
	PaymentPackageID string `gorm:"column:payment_package_id"`
	TransactionID    string `gorm:"column:transaction_id"`
	Name             string `gorm:"column:name"`
	Price            int64  `gorm:"column:price"`
	Credits          int64  `gorm:"column:credits"`
	Currency         string `gorm:"column:currency"`
	Category         string `gorm:"column:category"`
}

func toModel(dao transactionDAO) models.TransactionLedger {
	return models.TransactionLedger{
		ID:        dao.Id,
		WalletID:  dao.WalletID,
		UserID:    dao.UserID,
		Amount:    dao.Amount,
		Currency:  dao.Currency,
		Type:      dao.Type,
		Kind:      dao.Kind,
		IsReverse: dao.IsReverse,
		Status:    dao.Status,
		Provider:  dao.Provider,
		Metadata: models.TransactionLedgerMetadata{
			Token:                     dao.Token,
			DeviceID:                  dao.DeviceID,
			UserCouponID:              dao.UserCouponID,
			OriginalPrice:             dao.OriginalPrice,
			OriginalCredits:           dao.OriginalCredits,
			LineupID:                  dao.LineupID,
			ContestID:                 dao.ContestID,
			UserBalance:               dao.UserBalance,
			PromotionalBalance:        dao.PromotionalBalance,
			WinningsBalance:           dao.WinningsBalance,
			UserBalanceApplied:        dao.UserBalanceApplied,
			PromotionalBalanceApplied: dao.PromotionalBalanceApplied,
			WinningsBalanceApplied:    dao.WinningsBalanceApplied,
			Rake:                      dao.Rake,
			RakePercentage:            dao.RakePercentage,
			OriginalReferenceID:       dao.OriginalReferenceID,
			ReverseReason:             dao.ReverseReason,
			Clabe:                     dao.Clabe,
			BankAccountInformationID:  dao.BankAccountInformationID,
			Package: &models.PaymentPackage{
				ID:       dao.PaymentPackageDAO.PaymentPackageID,
				Name:     dao.PaymentPackageDAO.Name,
				Price:    dao.PaymentPackageDAO.Price,
				Credits:  dao.PaymentPackageDAO.Credits,
				Currency: dao.PaymentPackageDAO.Currency,
				Category: dao.PaymentPackageDAO.Category,
			},
		},
		CreatedAt: dao.CreatedAt,
		UpdatedAt: dao.UpdatedAt,
	}
}

func toDAO(model models.TransactionLedger) *transactionDAO {
	return &transactionDAO{
		Id:                        model.ID,
		WalletID:                  model.WalletID,
		UserID:                    model.UserID,
		Amount:                    model.Amount,
		Currency:                  model.Currency,
		Type:                      model.Type,
		Kind:                      model.Kind,
		IsReverse:                 model.IsReverse,
		Status:                    model.Status,
		Provider:                  model.Provider,
		Token:                     model.Metadata.Token,
		DeviceID:                  model.Metadata.DeviceID,
		UserCouponID:              model.Metadata.UserCouponID,
		OriginalPrice:             model.Metadata.OriginalPrice,
		OriginalCredits:           model.Metadata.OriginalCredits,
		LineupID:                  model.Metadata.LineupID,
		ContestID:                 model.Metadata.ContestID,
		UserBalance:               model.Metadata.UserBalance,
		PromotionalBalance:        model.Metadata.PromotionalBalance,
		WinningsBalance:           model.Metadata.WinningsBalance,
		UserBalanceApplied:        model.Metadata.UserBalanceApplied,
		PromotionalBalanceApplied: model.Metadata.PromotionalBalanceApplied,
		WinningsBalanceApplied:    model.Metadata.WinningsBalanceApplied,
		Rake:                      model.Metadata.Rake,
		RakePercentage:            model.Metadata.RakePercentage,
		OriginalReferenceID:       model.Metadata.OriginalReferenceID,
		ReverseReason:             model.Metadata.ReverseReason,
		Clabe:                     model.Metadata.Clabe,
		BankAccountInformationID:  model.Metadata.BankAccountInformationID,
		CreatedAt:                 model.CreatedAt,
		PaymentPackageDAO: PaymentPackageDAO{
			PaymentPackageID: model.Metadata.Package.ID,
			Name:             model.Metadata.Package.Name,
			Price:            model.Metadata.Package.Price,
			Credits:          model.Metadata.Package.Credits,
			Currency:         model.Metadata.Package.Currency,
			Category:         model.Metadata.Package.Category,
		},
	}
}

// TableName overrides the table name used by User to `profiles`
func (d transactionDAO) TableName() string {
	return "transaction"
}

func (p PaymentPackageDAO) TableName() string {
	return "payment_package"
}
