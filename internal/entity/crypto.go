package entity

type CryproCurrency struct {
	Id     int    `json:"id" db:"id"`
	AccountID int `json:"account_id" db:""`
	Coin   string `json:"coin" db:"coin"`
	Amount int    `json:"amount" db:"amount"`
}
