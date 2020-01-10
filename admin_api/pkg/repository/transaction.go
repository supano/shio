package repository

import (
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	adminApi "github.com/digithun/shio/admin_api"
)

type TransactionRepository interface {
	FindTransactionList(ctx context.Context) ([]*TransactionTest, error)
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

type TransactionTest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AssetId              string   `protobuf:"bytes,2,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	Status               int64   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Price                int64    `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	UserId               string   `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DescribeUrls         []string `protobuf:"bytes,6,rep,name=describe_urls,json=describeUrls,proto3" json:"describe_urls,omitempty"`
}

func (r *DefaultTransactionRepository) FindTransactionList(ctx context.Context) ([]*TransactionTest, error) {
	q := r.DefaultQuery()
	q.Offset(0)
	q.Limit(10)

	var dest []*TransactionTest
	_, err := r.client.GetAll(ctx, q, &dest)
	fmt.Println(dest)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}


	return dest, nil
}
