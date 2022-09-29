package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm-with-generics/pkg/models"
	"time"
)

type CouponApplicationDAO struct {
	ID             int64                        `gorm:"primaryKey,column:id"`
	TransactionID  string                       `gorm:"column:transaction_id"`
	CouponID       string                       `gorm:"column:coupon_id"`
	Name           string                       `gorm:"column:name"`
	Description    string                       `gorm:"column:description"`
	Code           string                       `gorm:"column:code"`
	Currency       string                       `gorm:"column:currency"`
	Capacity       int64                        `gorm:"column:capacity"`
	TimesPerUser   int64                        `gorm:"column:times_per_user"`
	Type           string                       `gorm:"column:type"`
	ValidUntil     time.Time                    `gorm:"column:valid_until"`
	ValidFrom      time.Time                    `gorm:"column:valid_from"`
	MinValidAmount int64                        `gorm:"column:min_valid_amount"`
	MaxValidAmount int64                        `gorm:"column:max_valid_amount"`
	MetadataDAO    CouponApplicationMetadataDAO `gorm:"embedded;embeddedPrefix:metadata_"`
	CreatedAt      time.Time                    `gorm:"column:created_at"`
}

type CouponApplicationCashRulesDAO struct {
	Amount int64  `json:"amount"`
	Type   string `json:"type"`
	Change string `json:"change"`
}

func (cs CouponApplicationCashRulesDAO) Value() (driver.Value, error) {
	return json.Marshal(cs)
}

// Scan Unmarshal
func (cs *CouponApplicationCashRulesDAO) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("cannot read bytes")
	}
	return json.Unmarshal(b, &cs)
}

type CouponApplicationMetadataDAO struct {
	ReferrerID string `gorm:"column:referrer_id"`
}

func toCouponApplicationDAO(m models.TransactionLedger) *CouponApplicationDAO {
	if m.Metadata.Coupon == nil {
		return nil
	}

	c := *m.Metadata.Coupon

	return &CouponApplicationDAO{
		ID:             time.Time.UnixNano(time.Now()),
		TransactionID:  m.ID,
		CouponID:       c.ID.Hex(),
		Name:           c.Name,
		Description:    c.Description,
		Code:           c.Code,
		Currency:       c.Currency,
		Capacity:       c.Capacity,
		TimesPerUser:   c.TimesPerUser,
		Type:           c.Type,
		ValidUntil:     c.ValidUntil.Time(),
		ValidFrom:      c.ValidFrom.Time(),
		MinValidAmount: c.MinValidAmount,
		MaxValidAmount: c.MaxValidAmount,
		MetadataDAO:    toCouponApplicationMetadataDAO(c),
	}
}

func toCouponApplicationMetadataDAO(m models.Coupon) CouponApplicationMetadataDAO {
	return CouponApplicationMetadataDAO{
		ReferrerID: m.Metadata.ReferrerID,
	}
}

func toCouponModel(ca CouponApplicationDAO) *models.Coupon {
	couponID, _ := primitive.ObjectIDFromHex(ca.CouponID)

	validUntil := primitive.NewDateTimeFromTime(ca.ValidUntil)
	validFrom := primitive.NewDateTimeFromTime(ca.ValidFrom)
	createdAt := primitive.NewDateTimeFromTime(ca.CreatedAt)

	return &models.Coupon{
		ID:             &couponID,
		Capacity:       ca.Capacity,
		TimesPerUser:   ca.TimesPerUser,
		Type:           ca.Type,
		ValidUntil:     &validUntil,
		ValidFrom:      &validFrom,
		MinValidAmount: ca.MinValidAmount,
		MaxValidAmount: ca.MaxValidAmount,
		Name:           ca.Name,
		Description:    ca.Description,
		Code:           ca.Code,
		Currency:       ca.Currency,
		CreatedAt:      &createdAt,
		Metadata:       toCouponMetadataModel(ca),
	}
}

func toCouponMetadataModel(ca CouponApplicationDAO) models.CouponMetadata {
	return models.CouponMetadata{
		ReferrerID: ca.MetadataDAO.ReferrerID,
	}
}

func (ca CouponApplicationDAO) TableName() string {
	return "coupon_applications"
}
