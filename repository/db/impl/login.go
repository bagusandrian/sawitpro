package impl

import (
	"context"
	"database/sql"
	"errors"

	"github.com/bagusandrian/sawitpro/helper"
	"github.com/bagusandrian/sawitpro/model"
)

func (r *repository) Login(ctx context.Context, req model.RequestLogin) (res model.ResponseLogin, err error) {
	var (
		password string
		id       int64
	)
	err = r.dbSlave.QueryRow("SELECT id, password FROM users WHERE phonenumber=$1", req.PhoneNumber).Scan(&id, &password)
	// validation phone number not register
	if err == sql.ErrNoRows {
		return res, errors.New("phone number is not register")
	}
	// validation err from server
	if err != nil {
		return res, err
	}
	// validation
	if !r.bcrypt.ComparePassword(password, req.Password) {
		return res, errors.New("password is wrong")
	}
	// processing to increment of total login and update time for login
	tx, err := r.dbMaster.Begin()
	defer tx.Rollback()
	_, err = tx.Exec("UPDATE users SET success_login = success_login+1 WHERE id =$1", id)
	if err != nil {
		return res, err
	}
	err = tx.Commit()
	if err != nil {
		return res, err
	}
	// Generate JWT
	token, err := helper.GenerateJWT(id, r.cfg.Server.JWTSecretKey)
	return model.ResponseLogin{
		ID:       id,
		JWTToken: token,
	}, err
}
