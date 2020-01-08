package endpointv1

import (
	"context"
	entityv1 "github.com/digithun/shio/admin/pkg/entity/v1"
	"time"
)

type AdminEndPoint struct{}

func (a *AdminEndPoint) Login(ctx context.Context,req *LoginRequest) (*LoginResponse, error) {
	session := entityv1.Session{
		Id:                   "success",
		UserId:               "success",
		CreatedAt:            time.Time.Unix(time.Now()),
	}
	return &LoginResponse{
		Session:              &session,
	}, nil
}

func NewAdminEndPoint() *AdminEndPoint  {
	return &AdminEndPoint{}
}