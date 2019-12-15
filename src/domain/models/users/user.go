package users

// UserID id
type UserID struct {
	ID int
}

// User user info
type User struct {
	Name  string
	Email Email
}

// Email email
type Email struct {
	Email string
}

// CreateUser create user
func (user *User) CreateUser() *UserID {
	userID := &UserID{1}

	return userID
}
