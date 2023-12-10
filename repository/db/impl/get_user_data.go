package impl

import (
	"context"

	"github.com/bagusandrian/sawitpro/model"
)

func (r *repository) GetUserDataByID(ctx context.Context, userID int64) (res model.ResponseGetProfile, err error) {
	var (
		fullName, phoneNumber string
	)
	err = r.dbSlave.QueryRow("SELECT phonenumber, fullname FROM users WHERE id = $1", userID).Scan(&phoneNumber, &fullName)
	if err != nil {
		return res, err
	}
	// Generate JWT
	return model.ResponseGetProfile{
		FullName:    fullName,
		PhoneNumber: phoneNumber,
	}, nil
}
