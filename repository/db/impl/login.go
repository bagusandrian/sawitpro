package impl

import (
	"context"
	"errors"

	"github.com/bagusandrian/sawitpro/helper"
	"github.com/bagusandrian/sawitpro/model"
	"golang.org/x/crypto/bcrypt"
)

func (r *repository) Login(ctx context.Context, req model.RequestLogin) (res model.ResponseLogin, err error) {
	var (
		password string
		id       int64
	)
	err = r.dbSlave.QueryRow("SELECT id, password FROM users WHERE phonenumber=$1", req.PhoneNumber).Scan(&id, &password)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(password), []byte(req.Password)) != nil {
		return res, errors.New("password is wrong")
	}
	// Generate JWT
	token, err := helper.GenerateJWT(id, r.cfg.Server.JWTSecretKey)
	return model.ResponseLogin{
		ID:       id,
		JWTToken: token,
	}, err
}
