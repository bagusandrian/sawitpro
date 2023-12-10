package impl

import (
	"context"
	"errors"

	"github.com/bagusandrian/sawitpro/model"
)

func (r *repository) Registration(ctx context.Context, req model.RequestRegistration) (res model.ResponseRegristration, err error) {
	tx, err := r.dbMaster.Begin()
	defer tx.Rollback()
	if err != nil {
		return res, err
	}
	result, err := tx.ExecContext(ctx, "INSERT INTO users (fullname,phonenumber,password) values ($1, $2, $3) ON CONFLICT DO NOTHING",
		req.FullName, req.PhoneNumber, r.bcrypt.GeneratePassword(req.Password))
	if err != nil {
		return res, err
	}
	rowChanges, _ := result.RowsAffected()
	if rowChanges == 0 {
		return res, errors.New("duplicate phone number")
	}
	tx.Commit()
	res = model.ResponseRegristration{
		FullName:    req.FullName,
		PhoneNumber: req.PhoneNumber,
	}
	return res, nil
}
