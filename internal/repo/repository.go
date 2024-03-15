package repo

import (
	"currency-operations/internal/entity"
	postgresdb "currency-operations/internal/repo/postgres_db"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GetUser(username, password string) (entity.User, error)
}

type Account interface {
	Create(userId int, account entity.Account) (int, error)
	GetById(userId, account_id int) (entity.Account, error)
	GetAll(userId int) ([]entity.Account, error)
	Deposit(userId, accountId, amount int) error
	Withdraw(userId, accountId, amount int) error
}

type Crypto interface {
	Create(accountId int, crypto entity.CryproCurrency) (int, error)
	GetById(userId, cryptoId int) (entity.CryproCurrency, error)
	GetAll(userId, accountId int) ([]entity.CryproCurrency, error)
	Buy(userId, cryptoId, amount int) error
}

type Repository struct {
	Authorization
	Account
	Crypto
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: postgresdb.NewAuthPostgres(db),
		Account:       postgresdb.NewAccountPostgres(db),
		Crypto:        postgresdb.NewCryptoPostgres(db),
	}
}
