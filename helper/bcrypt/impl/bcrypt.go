package impl

import "golang.org/x/crypto/bcrypt"

func (r *repository) GeneratePassword(password string) string {
	bytePassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytePassword)
}
func (r *repository) ComparePassword(passwordReq, PasswordDB string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordReq), []byte(PasswordDB))
	return err == nil
}
