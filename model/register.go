package model

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/bagusandrian/sawitpro/helper"
)

type (
	RequestRegistration struct {
		FullName    string
		PhoneNumber string
		Password    string
	}
	ResponseRegristration struct {
		ID           int64    `json:"id"`
		FullName     string   `json:"fullname"`
		PhoneNumber  string   `json:"phonenumber"`
		ErrorMessage []string `json:"error_message"`
	}
)

var rgxSQLInjectorChar = regexp.MustCompile(`[\'\"\t\r\\\n;]`)

func (model *RequestRegistration) ParseFromHTTPRequest(r *http.Request) []string {
	var fullName, phoneNumber, password string
	fullName = rgxSQLInjectorChar.ReplaceAllString(r.FormValue("fullname"), " ")
	phoneNumber = rgxSQLInjectorChar.ReplaceAllString(r.FormValue("phonenumber"), " ")
	password = rgxSQLInjectorChar.ReplaceAllString(r.FormValue("password"), " ")
	errors := []string{}
	if fullName == "" || phoneNumber == "" || password == "" {
		errors = append(errors, "Data is empty")
	}
	if !(len(phoneNumber) >= 10 && len(phoneNumber) <= 13 && strings.HasPrefix(phoneNumber, "+62")) {
		errors = append(errors, "Invalid phone number")
	}
	if !(len(fullName) >= 3 && len(fullName) <= 60) {
		errors = append(errors, "Invalid full name")
	}
	if !(len(password) >= 6 && len(password) <= 64 &&
		helper.ContainsUppercase(password) && helper.ContainsDigit(password) && helper.ContainsSpecialChar(password)) {
		errors = append(errors, "Invalid password")
	}
	if len(errors) > 0 {
		return errors
	}
	model.FullName = fullName
	model.PhoneNumber = phoneNumber
	model.Password = password

	return nil
}
