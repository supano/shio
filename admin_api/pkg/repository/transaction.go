package repository

import (
	"cloud.google.com/go/datastore"
	"context"
	adminApi "github.com/digithun/shio/admin_api"
	entityv1 "github.com/digithun/shio/admin_api/pkg/entity/v1"
)

type TransactionRepository interface {
	FindTransactionList(ctx context.Context) ([]*entityv1.Transaction, error)
}

type DefaultTransactionRepository struct {
	client *datastore.Client
}

func NewDefaultTransactionRepository(client *datastore.Client) *DefaultTransactionRepository  {
	return &DefaultTransactionRepository{client: client}
}

func (r *DefaultTransactionRepository) DefaultQuery() *datastore.Query  {
	return datastore.NewQuery(adminApi.TransactionKind).Namespace(adminApi.ShioNameSpace)
}

func (r *DefaultTransactionRepository) FindTransactionList(ctx context.Context) ([]*entityv1.Transaction, error) {
	q := r.DefaultQuery()
	q.Offset(0)
	q.Limit(10)

	var dest []*entityv1.Transaction
	_, err := r.client.GetAll(ctx, q, &dest)
	if err != nil {
		return nil, err
	}

	return dest, nil
}
