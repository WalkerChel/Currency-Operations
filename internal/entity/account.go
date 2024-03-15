package entity

import "errors"

type Account struct {
	Id       int    `json:"id" db:"id"`
	OwnerId  int    `json:"owner_id" db:"owner_id"`
	Currency string `json:"currency" db:"currency"`
	Balance  int    `json:"balance" db:"balance"`
}

// it's more reasonable to make handler for withdrawing
// type UpdateAccount struct {
// 	Balance  int    `json:"balance"`
// }

type UpdateAccount struct {
	Balance *int `json:"balance"`
}

func (a *UpdateAccount) Validate() error {
	if a.Balance == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
