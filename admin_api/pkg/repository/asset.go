package repository

import (
	"cloud.google.com/go/datastore"
	"context"
	adminApi "github.com/digithun/shio/admin_api"
	entityv1 "github.com/digithun/shio/admin_api/pkg/entity/v1"
)

type AssetRepository interface {
	FindAssetList(ctx context.Context) ([]*entityv1.Asset, error)
}

type DefaultAssetRepository struct {
	client *datastore.Client
}

func NewDefaultAssetRepository(client *datastore.Client) *DefaultAssetRepository  {
	return &DefaultAssetRepository{client: client}
}

func (r *DefaultAssetRepository) DefaultQuery() *datastore.Query  {
	return datastore.NewQuery(adminApi.AssetKind).Namespace(adminApi.ShioNameSpace)
}


func (r *DefaultAssetRepository) FindAssetList(ctx context.Context) ([]*entityv1.Asset, error)  {
	q := r.DefaultQuery()
	q.Offset(0)
	q.Limit(10)

	var dest []*entityv1.Asset
	_, err := r.client.GetAll(ctx, q, &dest)
	if err != nil {
		return nil, err
	}

	return dest, nil
}