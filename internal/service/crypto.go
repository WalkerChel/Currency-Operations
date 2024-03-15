package service

import (
	"currency-operations/internal/entity"
	"currency-operations/internal/repo"
)

type CryptoService struct {
	repo        repo.Crypto
	accountRepo repo.Account
}

func NewCryptoService(repo repo.Crypto, accountRepo repo.Account) *CryptoService {
	return &CryptoService{
		repo:        repo,
		accountRepo: accountRepo,
	}
}

func (cs *CryptoService) Create(userId, accountId int, crypto entity.CryproCurrency) (int, error) {
	_, err := cs.accountRepo.GetById(userId, accountId)
	if err != nil {
		//account does't exist or does't belong to user
		return 0, err
	}
	return cs.repo.Create(accountId, crypto)
}

func (cs *CryptoService) GetById(userId, cryptoId int) (entity.CryproCurrency, error) {
	return entity.CryproCurrency{}, nil
}

func (cs *CryptoService) GetAll(userId, accountId int) ([]entity.CryproCurrency, error) {
	return nil, nil
}

func (cs *CryptoService) Buy(userId, cryptoId, amount int) error {
	return nil
}
