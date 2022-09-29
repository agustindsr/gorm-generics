package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Coupon struct {
	ID             *primitive.ObjectID `json:"id"                    bson:"_id,omitempty"`
	CashRules      CashRules           `json:"cashRules"             bson:"cashRules,omitempty"`
	Capacity       int64               `json:"capacity,omitempty"    bson:"capacity,omitempty"` // if capacity == 0, that means capacity is infinite.
	TimesPerUser   int64               `json:"timesPerUser"          bson:"timesPerUser,omitempty"`
	Type           string              `json:"type"                  bson:"type,omitempty"`
	ValidUntil     *primitive.DateTime `json:"validUntil,omitempty"  bson:"validUntil,omitempty"`
	ValidFrom      *primitive.DateTime `json:"validFrom,omitempty"   bson:"validFrom,omitempty"`
	MinValidAmount int64               `json:"minValidAmount"        bson:"minValidAmount"` // minimum price of package to apply coupon
	MaxValidAmount int64               `json:"maxValidAmount"        bson:"maxValidAmount"` // maximum value coupon will apply to (only for percentage coupons). If MaxValidAmount = 0, that means maxValidAmount if infinite.
	Name           string              `json:"name,omitempty"        bson:"name,omitempty"`
	Description    string              `json:"description,omitempty" bson:"description,omitempty"`
	Code           string              `json:"code,omitempty"        bson:"code,omitempty"`
	Currency       string              `json:"currency,omitempty"    bson:"currency,omitempty"`
	CreatedAt      *primitive.DateTime `json:"createdAt,omitempty"   bson:"createdAt,omitempty"`
	UpdatedAt      *primitive.DateTime `json:"updatedAt,omitempty"   bson:"updatedAt,omitempty"`
	Metadata       CouponMetadata      `json:"metadata"              bson:"metadata"`
}

type CashRules struct {
	Amount int64  `json:"amount,omitempty" bson:"amount,omitempty"`
	Type   string `json:"type,omitempty"   bson:"type,omitempty"`
	Change string `json:"change,omitempty" bson:"change,omitempty"`
}

type CouponMetadata struct {
	ReferrerID string `json:"referrerId" bson:"referrerId,omitempty"`
}
