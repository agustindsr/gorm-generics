package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentPackage struct {
	ID              *primitive.ObjectID `json:"id"              ion:"id"              bson:"_id,omitempty"`
	Name            string              `json:"name"            ion:"name"            bson:"name,omitempty"`
	Price           int64               `json:"price"           ion:"price"           bson:"price,omitempty"`
	Credits         int64               `json:"credits"         ion:"credits"         bson:"credits,omitempty"`
	Sort            int                 `json:"sort"            ion:"sort"            bson:"sort,omitempty"`
	Currency        string              `json:"currency"        ion:"currency"        bson:"currency,omitempty"`
	Image           string              `json:"image"           ion:"image"           bson:"image,omitempty"`
	ImageWithCoupon string              `json:"imageWithCoupon" ion:"imageWithCoupon" bson:"imageWithCoupon,omitempty"`
	Category        string              `json:"category"        ion:"category"        bson:"category,omitempty"`
	Active          bool                `json:"active"          ion:"active"          bson:"active,omitempty"`
	CreatedAt       *primitive.DateTime `json:"createdAt"       ion:"createdAt"       bson:"createdAt,omitempty"`
	UpdatedAt       *primitive.DateTime `json:"updatedAt"       ion:"updatedAt"       bson:"updatedAt,omitempty"`
}
