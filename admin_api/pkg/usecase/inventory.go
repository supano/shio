package usecase

import (
	"context"
	entityv1 "github.com/digithun/shio/admin_api/pkg/entity/v1"
	"github.com/digithun/shio/admin_api/pkg/repository"
)

type InventoryUseCase interface {
	GetAssetList(ctx context.Context) ([]*entityv1.Asset, error)
}

type DefaultInventoryUseCase struct {
	asset repository.AssetRepository
}

func NewDefaultInventoryUseCase(asset repository.AssetRepository) *DefaultInventoryUseCase {
	return &DefaultInventoryUseCase{
		asset: asset,
	}
}

func (u *DefaultInventoryUseCase) GetAssetList(ctx context.Context) ([]*entityv1.Asset, error) {
	return u.asset.FindAssetList(ctx)
}
