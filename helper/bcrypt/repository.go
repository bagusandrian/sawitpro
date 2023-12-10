package bcrypt

//go:generate mockery --name=Repository --filename=mock_Repository.go --inpackage
type Repository interface {
	GeneratePassword(password string) string
	ComparePassword(password, passwordHash string) bool
}
