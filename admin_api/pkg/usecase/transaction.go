package usecase

import (
	"context"
	entityv1 "github.com/digithun/shio/admin_api/pkg/entity/v1"
	"github.com/digithun/shio/admin_api/pkg/repository"
)

type TransactionUseCase interface {
	GetTransactionList(ctx context.Context) ([]*entityv1.Transaction, error)
}

type DefaultTransactionUseCase struct {
	transaction repository.TransactionRepository
}

func NewDefaultTransactionUseCase(transaction repository.TransactionRepository) *DefaultTransactionUseCase {
	return &DefaultTransactionUseCase{
		transaction: transaction,
	}
}

func (u *DefaultTransactionUseCase) GetTransactionList(ctx context.Context) ([]*entityv1.Transaction, error) {
	return u.transaction.FindTransactionList(ctx)
}

