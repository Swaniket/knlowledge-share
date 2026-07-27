package types

type User struct {
	ID       string // User mongo object ID
	Username string
	Email    string
	Password string
	Role     string // Possible Values: "learner" / "teacher" / "admin"
}
