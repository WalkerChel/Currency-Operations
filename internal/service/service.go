package service

import (
	"currency-operations/internal/entity"
	"currency-operations/internal/repo"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Account interface {
	Create(id int, account entity.Account) (int, error)
	GetById(userId, account_id int) (entity.Account, error)
	GetAll(userId int) ([]entity.Account, error)
	Deposit(userId, accountId, amount int) error
	Withdraw(userId, accountId, amount int) error
}

type Crypto interface {
	Create(userId, accountId int, crypto entity.CryproCurrency) (int, error)
	GetById(userId, cryptoId int) (entity.CryproCurrency, error)
	GetAll(userId, accountId int) ([]entity.CryproCurrency, error)
	Buy(userId, cryptoId, amount int) error
}

type Service struct {
	Authorization
	Account
	Crypto
}

func NewService(repos *repo.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Account:       NewAccountService(repos.Account),
		Crypto:        NewCryptoService(repos.Crypto, repos.Account),
	}
}
