package usecase

import (
	"context"
	"github.com/digithun/shio/admin_api/pkg/repository"
)

type TransactionUseCase interface {
	GetTransactionList(ctx context.Context) ([]*repository.TransactionTest, error)
}

type DefaultTransactionUseCase struct {
	transaction repository.TransactionRepository
}

func NewDefaultTransactionUseCase(transaction repository.TransactionRepository) *DefaultTransactionUseCase {
	return &DefaultTransactionUseCase{
		transaction: transaction,
	}
}

func (u *DefaultTransactionUseCase) GetTransactionList(ctx context.Context) ([]*repository.TransactionTest, error) {
	return u.transaction.FindTransactionList(ctx)
}

