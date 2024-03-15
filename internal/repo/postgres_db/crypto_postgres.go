package postgresdb

import (
	"currency-operations/internal/entity"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type CryptoPostgres struct {
	db *sqlx.DB
}

func NewCryptoPostgres(db *sqlx.DB) *CryptoPostgres {
	return &CryptoPostgres{db: db}
}

func (cp *CryptoPostgres) Create(accountId int, crypto entity.CryproCurrency) (int, error) {

	var id int

	createCryptoQuery := fmt.Sprintf(`INSERT INTO %s (account_id, coin) values ($1, $2) RETURNING id`,
		cryptoTable)

	row := cp.db.QueryRow(createCryptoQuery, accountId, crypto.Coin)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (cp *CryptoPostgres) GetById(userId, cryptoId int) (entity.CryproCurrency, error) {
	return entity.CryproCurrency{}, nil
}

func (cp *CryptoPostgres) GetAll(userId, accountId int) ([]entity.CryproCurrency, error) {
	return nil, nil
}

func (cp *CryptoPostgres) Buy(userId, cryptoId, amount int) error {
	return nil
}
