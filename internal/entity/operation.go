package entity

import (
	"errors"
	"time"
)

type ExchangeOperation struct {
	Id            int       `json:"id" db:"id"`
	Currency      string    `json:"currrency" db:"currency"`
	Amount        int       `json:"amount" db:"amount"`
	Coin          string    `json:"coin" db:"coin"`
	PurchasePrise int       `json:"id" db:"purch_prise"`
	DesireTime    time.Time `json:"desire_time" db:"desire_time"`
	Status        bool      `json:"status" db:"status"`
}

type UpdateOperation struct {
	Amount        *int       `json:"amount" db:"amount"`
	Coin          *string    `json:"coin" db:"coin"`
	PurchasePrise *int       `json:"id" db:"purch_prise"`
	DesireTime    *time.Time `json:"desire_time" db:"desire_time"`
	Status        *bool      `json:"status" db:"status"`
}

func (r *UpdateOperation) Validate() error {
	if r.Amount == nil && r.Coin == nil && r.DesireTime == nil && r.Status == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
