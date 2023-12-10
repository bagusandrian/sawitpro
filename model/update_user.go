package model

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/bagusandrian/sawitpro/helper"
)

type (
	RequestUpdateProfile struct {
		ID          int64
		FullName    string
		PhoneNumber string
	}
	ResponseUpdateProfile struct {
		FullName       string
		PhoneNumber    string
		SuccessMessage string
		ErrorMessage   string
	}
)

func (model *RequestUpdateProfile) ParseFromHTTPRequest(r *http.Request, secretKey []byte) error {
	var phoneNumber, fullName string
	phoneNumber = rgxSQLInjectorChar.ReplaceAllString(r.FormValue("phonenumber"), " ")
	fullName = rgxSQLInjectorChar.ReplaceAllString(r.FormValue("fullname"), " ")
	if phoneNumber == "" || fullName == "" {
		return errors.New("data empty")
	}
	if !(len(phoneNumber) >= 10 && len(phoneNumber) <= 13 && strings.HasPrefix(phoneNumber, "+62")) {
		return errors.New("invalid phone number")
	}
	if !(len(fullName) >= 3 && len(fullName) <= 60) {
		return errors.New("invalid full name")
	}
	tokenString := helper.ExtractTokenFromHeader(r)
	claims, err := helper.VerifyToken(tokenString, secretKey)
	if err != nil {
		return errors.New("token invalid")
	}
	userIDstr, err := claims.GetIssuer()
	if err != nil {
		return errors.New("user not found")
	}
	userID, err := strconv.ParseInt(userIDstr, 10, 64)
	if err != nil {
		return errors.New("user not found")
	}

	model.ID = userID
	model.PhoneNumber = phoneNumber
	model.FullName = fullName

	return nil
}
