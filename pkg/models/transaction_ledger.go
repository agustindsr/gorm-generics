package models

import "time"

type TransactionLedger struct {
	ID        string                    `json:"id"`
	WalletID  string                    `json:"walletId"`
	UserID    string                    `json:"userId"`
	Amount    int64                     `json:"amount"`
	Currency  string                    `json:"currency"`
	Type      string                    `json:"type"`
	Kind      string                    `json:"kind"`
	IsReverse bool                      `json:"isReverse"`
	Status    string                    `json:"status"`
	Provider  string                    `json:"provider"`
	Metadata  TransactionLedgerMetadata `json:"metadata"`
	CreatedAt time.Time                 `json:"createdAt"`
	UpdatedAt time.Time                 `json:"updatedAt"`
}

type TransactionLedgerMetadata struct {
	Package                   *PaymentPackage `json:"package"`
	Coupon                    *Coupon         `json:"coupon"`
	Token                     string          `json:"token"`
	DeviceID                  string          `json:"deviceId"`
	UserCouponID              string          `json:"userCouponId"`
	OriginalPrice             int64           `json:"originalPrice"`
	OriginalCredits           int64           `json:"originalCredits"`
	LineupID                  string          `json:"lineupId"`
	ContestID                 string          `json:"contestId"`
	UserBalance               int64           `json:"userBalance"`
	PromotionalBalance        int64           `json:"promotionalBalance"`
	WinningsBalance           int64           `json:"winningsBalance"`
	UserBalanceApplied        int64           `json:"userBalanceApplied"`
	PromotionalBalanceApplied int64           `json:"promotionalBalanceApplied"`
	WinningsBalanceApplied    int64           `json:"winningsBalanceApplied"`
	Rake                      int64           `json:"rake"`
	RakePercentage            float64         `json:"rakePercentage"`
	OriginalReferenceID       string          `json:"originalReferenceId"`
	ReverseReason             string          `json:"reverseReason"`
	Clabe                     string          `json:"clabe"`
	BankAccountInformationID  string          `json:"bankAccountInformationId"`
}
