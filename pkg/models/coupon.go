package models

import "time"

type Coupon struct {
	ID             string         `json:"id"`
	CashRules      CashRules      `json:"cashRules"`
	Capacity       int64          `json:"capacity,omitempty"`
	TimesPerUser   int64          `json:"timesPerUser"`
	Type           string         `json:"type"`
	ValidUntil     time.Time      `json:"validUntil,omitempty"`
	ValidFrom      time.Time      `json:"validFrom,omitempty"`
	MinValidAmount int64          `json:"minValidAmount"`
	MaxValidAmount int64          `json:"maxValidAmount"`
	Name           string         `json:"name,omitempty"`
	Description    string         `json:"description,omitempty"`
	Code           string         `json:"code,omitempty"`
	Currency       string         `json:"currency,omitempty"`
	CreatedAt      time.Time      `json:"createdAt,omitempty"`
	UpdatedAt      time.Time      `json:"updatedAt,omitempty"`
	Metadata       CouponMetadata `json:"metadata"`
}

type CashRules struct {
	Amount int64  `json:"amount,omitempty"`
	Type   string `json:"type,omitempty"`
	Change string `json:"change,omitempty"`
}

type CouponMetadata struct {
	ReferrerID string `json:"referrerId"`
}
