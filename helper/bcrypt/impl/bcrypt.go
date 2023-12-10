package impl

import "golang.org/x/crypto/bcrypt"

func (r *repository) GeneratePassword(password string) string {
	bytePassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytePassword)
}
func (r *repository) ComparePassword(passwordHash, passwordReq string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(passwordReq))
	return err == nil
}
