package postgresdb

import (
	"currency-operations/internal/entity"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AccountPostgres struct {
	db *sqlx.DB
}

func NewAccountPostgres(db *sqlx.DB) *AccountPostgres {
	return &AccountPostgres{
		db: db,
	}
}

func (a *AccountPostgres) Create(userId int, account entity.Account) (int, error) {
	var currency string

	getCurrencytQuery := fmt.Sprintf(`SELECT ac.currency FROM %s ac
										WHERE ac.owner_id = $1 AND ac.currency = $2`,
		accountTable)

	err := a.db.Get(&currency, getCurrencytQuery, userId, account.Currency)
	if err == nil {
		return 0, fmt.Errorf("account with currency %s already exists", account.Currency)
	}

	var id int

	createAccountQuery := fmt.Sprintf("INSERT INTO %s (owner_id, currency) VALUES (($1), $2) RETURNING ID", accountTable)

	row := a.db.QueryRow(createAccountQuery, userId, account.Currency)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (a *AccountPostgres) GetById(userId, account_id int) (entity.Account, error) {
	var account entity.Account

	query := fmt.Sprintf(`SELECT * FROM %s ac
							WHERE ac.owner_id = $1 AND ac.id = $2`,
		accountTable)

	err := a.db.Get(&account, query, userId, account_id)

	return account, err
}

func (a *AccountPostgres) GetAll(userId int) ([]entity.Account, error) {
	var accounts []entity.Account

	query := fmt.Sprintf(`SELECT * FROM %s ac
							WHERE ac.owner_id = $1`,
		accountTable)

	err := a.db.Select(&accounts, query, userId)

	return accounts, err
}

func (a *AccountPostgres) Deposit(userId, accountId, amount int) error {
	query := fmt.Sprintf(`UPDATE %s ac
							SET balance = balance + $1
							WHERE ac.owner_id = $2 AND ac.id = $3`,
		accountTable)

	_, err := a.db.Exec(query, amount, userId, accountId)
	return err
}

func (a *AccountPostgres) Withdraw(userId, accountId, amount int) error {
	query := fmt.Sprintf(`UPDATE %s ac
	SET balance = balance - $1
	WHERE ac.owner_id = $2 AND ac.id = $3`,
		accountTable)

	_, err := a.db.Exec(query, amount, userId, accountId)
	return err
}
