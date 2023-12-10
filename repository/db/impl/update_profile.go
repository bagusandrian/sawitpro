package impl

import (
	"context"

	"github.com/bagusandrian/sawitpro/model"
)

func (r *repository) UpdateProfile(ctx context.Context, req model.RequestUpdateProfile) (res model.ResponseUpdateProfile, err error) {
	tx, err := r.dbMaster.Begin()
	defer tx.Rollback()
	if err != nil {
		return res, err
	}
	// #1 update profile name
	_, err = tx.ExecContext(ctx, "UPDATE users set fullname=$1, update_time = now() WHERE id = $2",
		req.FullName, req.ID)
	if err != nil {
		return res, err
	}
	// #2 update phone number
	_, err = tx.ExecContext(ctx, "UPDATE users set phonenumber=$1, update_time = now() WHERE id = $2",
		req.PhoneNumber, req.ID)
	if err != nil {
		return res, err
	}
	tx.Commit()
	res = model.ResponseUpdateProfile{
		SuccessMessage: "success",
		FullName:       req.FullName,
		PhoneNumber:    req.PhoneNumber,
	}
	return
}
