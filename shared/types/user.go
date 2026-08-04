package types

type UserRole string

const (
	RoleLearner UserRole = "learner"
	RoleCreator UserRole = "creator"
	RoleAdmin   UserRole = "admin"
)

type User struct {
	ID       string // User mongo object ID
	Username string
	Email    string
	Password string
	Role     UserRole
}
