package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)
type Invoice struct {
	ID             bson.ObjectID `bson:"_id,omitempty" json:"id"`
	InvoiceID      string        `bson:"invoice_id" json:"invoice_id"`
	OrderID        string        `bson:"order_id" json:"order_id"`

	PaymentMethod  *string       `bson:"payment_method" json:"payment_method" validate:"required,oneof=CARD CASH"`
	PaymentStatus  *string       `bson:"payment_status" json:"payment_status" validate:"required,oneof=PENDING PAID"`

	PaymentDueDate time.Time     `bson:"payment_due_date" json:"payment_due_date"`

	CreatedAt      time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt      time.Time     `bson:"updated_at" json:"updated_at"`
}