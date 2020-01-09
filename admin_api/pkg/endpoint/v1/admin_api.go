package endpointv1

import (
	"context"
	entityv1 "github.com/digithun/shio/admin_api/pkg/entity/v1"
	"github.com/digithun/shio/admin_api/pkg/usecase"
	"time"
)

type AdminEndPoint struct{
	inventory usecase.InventoryUseCase
	transaction usecase.TransactionUseCase
}

func NewAdminEndPoint(fulfillment usecase.InventoryUseCase, transaction usecase.TransactionUseCase) *AdminEndPoint  {
	return &AdminEndPoint{
		inventory: fulfillment,
		transaction: transaction,
	}
}

func (e *AdminEndPoint) Login(ctx context.Context,req *LoginRequest) (*LoginResponse, error) {
	session := entityv1.Session{
		Id:                   "success",
		UserId:               "success",
		CreatedAt:            time.Time.Unix(time.Now()),
	}
	return &LoginResponse{
		Session:              &session,
	}, nil
}


func (e *AdminEndPoint) GetAssetList(ctx context.Context,req *GetAssetListRequest) (*GetAssetListResponse, error) {
	list, err := e.inventory.GetAssetList(ctx)
	if err != nil {
		return nil, err
	}

	return &GetAssetListResponse{
		List:            list,
	}, nil
}

func (e *AdminEndPoint) GetTransactionList(ctx context.Context,req *GetTransactionListRequest) (*GetTransactionListResponse, error)  {
	list, err := e.transaction.GetTransactionList(ctx)
	if err != nil {
		return nil, err
	}

	return &GetTransactionListResponse{
		List:            list,
	}, nil
}