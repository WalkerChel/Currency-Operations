package service

import (
	"currency-operations/internal/entity"
	"currency-operations/internal/repo"
	"errors"
)

type AccountService struct {
	repo repo.Account
}

func NewAccountService(repo repo.Account) *AccountService {
	return &AccountService{
		repo: repo,
	}
}

func (a *AccountService) Create(id int, account entity.Account) (int, error) {
	return a.repo.Create(id, account)
}

func (a *AccountService) GetAll(userId int) ([]entity.Account, error) {
	return a.repo.GetAll(userId)
}

func (a *AccountService) GetById(userId, account_id int) (entity.Account, error) {
	return a.repo.GetById(userId, account_id)
}

func (a *AccountService) Deposit(userId, accountId, amount int) error {
	if amount < 0 {
		return errors.New("can not deposit with negative amount")
	}
	return a.repo.Deposit(userId, accountId, amount)
}

func (a *AccountService) Withdraw(userId, accountId, amount int) error {
	if amount < 0 {
		return errors.New("withdraw amount can not be negative")
	}

	account, err := a.repo.GetById(userId, accountId)
	if err != nil {
		return err
	}

	if amount > account.Balance {
		return errors.New("not enough funds to withdraw")
	}

	return a.repo.Withdraw(userId, accountId, amount)
}
