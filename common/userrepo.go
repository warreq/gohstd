package common

const UserExistsError = `a user by that name already exists`

// UserRepo [sitory] is an interface wrapping the functions for working with
// user data
type UserRepo interface {
	// InsertUser commits a User to the storage, returning an error if the user
	// already exists
	InsertUser(user User, secret Secret) (err error)

	// GetUserByName queries the storage for a representation of a User, returning
	// nil and an error if the user does not exist
	GetUserByName(username string) (user User, err error)
}
