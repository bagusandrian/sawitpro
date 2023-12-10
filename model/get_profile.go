package model

type (
	// RequestGetProfile struct {
	// 	FullName    string
	// 	PhoneNumber string
	// 	Password    string
	// }
	ResponseGetProfile struct {
		FullName    string
		PhoneNumber string
	}
)

// func (model *RequestLogin) ParseFromHTTPRequest(r *http.Request) []string {
// 	var phoneNumber, password string
// 	phoneNumber = rgxSQLInjectorChar.ReplaceAllString(r.FormValue("phonenumber"), " ")
// 	password = rgxSQLInjectorChar.ReplaceAllString(r.FormValue("password"), " ")

// 	errors := []string{}
// 	if !(len(phoneNumber) >= 10 && len(phoneNumber) <= 13 && strings.HasPrefix(phoneNumber, "+62")) {
// 		errors = append(errors, "Invalid phone number")
// 	}
// 	if !(len(password) >= 6 && len(password) <= 64 &&
// 		helper.ContainsUppercase(password) && helper.ContainsDigit(password) && helper.ContainsSpecialChar(password)) {
// 		errors = append(errors, "Invalid password")
// 	}
// 	if len(errors) > 0 {
// 		return errors
// 	}
// 	model.PhoneNumber = phoneNumber
// 	model.Password = password

// 	return nil
// }
