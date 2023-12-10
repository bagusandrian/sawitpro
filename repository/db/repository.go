package db

import (
	"context"

	"github.com/bagusandrian/sawitpro/model"
)

//go:generate mockery --name=Repository --filename=mock_Repository.go --inpackage
type Repository interface {
	Registration(ctx context.Context, req model.RequestRegistration) (res model.ResponseRegristration, err error)
	Login(ctx context.Context, req model.RequestLogin) (res model.ResponseLogin, err error)
	GetUserDataByID(ctx context.Context, userID int64) (res model.ResponseGetProfile, err error)
	UpdateProfile(ctx context.Context, req model.RequestUpdateProfile) (res model.ResponseUpdateProfile, err error)
}
