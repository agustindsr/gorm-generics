package models

import (
	"gorm-with-generics/pkg/models"
	"time"
)

type DAO struct {
	ID                  string                `gorm:"column:id"`
	WalletID            string                `gorm:"column:wallet_id"`
	UserID              string                `gorm:"column:user_id"`
	DeviceID            string                `gorm:"column:device_id"`
	Amount              int64                 `gorm:"column:amount"`
	Currency            string                `gorm:"column:currency"`
	Type                string                `gorm:"column:type"`
	Kind                string                `gorm:"column:kind"`
	IsReverse           bool                  `gorm:"column:is_reverse"`
	ReverseReason       string                `gorm:"column:reverse_reason"`
	OriginalReferenceID string                `gorm:"column:original_reference_id"`
	Status              string                `gorm:"column:status"`
	Provider            string                `gorm:"column:provider"`
	OriginalPrice       int64                 `gorm:"column:original_price"`
	OriginalCredits     int64                 `gorm:"column:original_credits"`
	Metadata            MetadataDAO           `gorm:"embedded;embeddedPrefix:metadata_"`
	CreatedAt           time.Time             `gorm:"column:createdAt"`
	UpdatedAt           time.Time             `gorm:"column:updatedAt"`
	PackageApplication  PackageApplicationDAO `gorm:"foreignKey:TransactionID"`
	CouponApplication   CouponApplicationDAO  `gorm:"foreignKey:TransactionID"`
	BalanceApplication  BalanceApplicationDAO `gorm:"foreignKey:TransactionID"`
}

type MetadataDAO struct {
	Token                    string `gorm:"column:token"`
	UserCouponID             string `gorm:"column:user_coupon_id"`
	LineupID                 string `gorm:"column:lineup_id"`
	ContestID                string `gorm:"column:contest_id"`
	Clabe                    string `gorm:"column:clabe"`
	BankAccountInformationID string `gorm:"column:bank_account_information_id"`
}

func ToDAO(m models.TransactionLedger) DAO {
	return DAO{
		ID:                  m.ID,
		WalletID:            m.WalletID,
		UserID:              m.UserID,
		DeviceID:            m.Metadata.DeviceID,
		Amount:              m.Amount,
		Currency:            m.Currency,
		Type:                m.Type,
		Kind:                m.Kind,
		IsReverse:           m.IsReverse,
		ReverseReason:       m.Metadata.ReverseReason,
		OriginalReferenceID: m.Metadata.OriginalReferenceID,
		Status:              m.Status,
		Provider:            m.Provider,
		OriginalPrice:       m.Metadata.OriginalPrice,
		OriginalCredits:     m.Metadata.OriginalCredits,
		Metadata:            toMetadataDAO(m.Metadata),
		CreatedAt:           m.CreatedAt,
		UpdatedAt:           m.UpdatedAt,
		PackageApplication:  toPackageApplicationDAO(m),
		CouponApplication:   toCouponApplicationDAO(m),
		BalanceApplication:  toBalanceApplicationDAO(m),
	}
}

func toMetadataDAO(m models.TransactionLedgerMetadata) MetadataDAO {
	return MetadataDAO{
		Token:                    m.Token,
		UserCouponID:             m.UserCouponID,
		LineupID:                 m.LineupID,
		ContestID:                m.ContestID,
		Clabe:                    m.Clabe,
		BankAccountInformationID: m.BankAccountInformationID,
	}
}

func ToModel(d DAO) models.TransactionLedger {
	return models.TransactionLedger{
		ID:        d.ID,
		WalletID:  d.WalletID,
		UserID:    d.UserID,
		Amount:    d.Amount,
		Currency:  d.Currency,
		Type:      d.Type,
		Kind:      d.Kind,
		IsReverse: d.IsReverse,
		Status:    d.Status,
		Provider:  d.Provider,
		Metadata:  toMetadata(d),
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

func toMetadata(d DAO) models.TransactionLedgerMetadata {
	return models.TransactionLedgerMetadata{
		Package:                   toPackageModel(d.PackageApplication),
		Coupon:                    toCouponModel(d.CouponApplication),
		Token:                     d.Metadata.Token,
		DeviceID:                  d.DeviceID,
		OriginalPrice:             d.OriginalPrice,
		OriginalCredits:           d.OriginalCredits,
		LineupID:                  d.Metadata.LineupID,
		ContestID:                 d.Metadata.ContestID,
		UserCouponID:              d.Metadata.UserCouponID,
		UserBalance:               d.BalanceApplication.UserBalance,
		PromotionalBalance:        d.BalanceApplication.PromotionalBalance,
		WinningsBalance:           d.BalanceApplication.WinningsBalance,
		UserBalanceApplied:        d.BalanceApplication.UserBalanceApplied,
		PromotionalBalanceApplied: d.BalanceApplication.PromotionalBalanceApplied,
		WinningsBalanceApplied:    d.BalanceApplication.WinningsBalanceApplied,
		OriginalReferenceID:       d.OriginalReferenceID,
		ReverseReason:             d.ReverseReason,
		Clabe:                     d.Metadata.Clabe,
		BankAccountInformationID:  d.Metadata.BankAccountInformationID,
	}
}

func ToModels(dao []DAO) []models.TransactionLedger {
	transactions := make([]models.TransactionLedger, 0)
	if len(dao) == 0 {
		return transactions
	}

	for _, d := range dao {
		transaction := ToModel(d)

		transactions = append(transactions, transaction)
	}

	return transactions
}

func (d DAO) Table() string {
	return "transaction_ledger"
}
