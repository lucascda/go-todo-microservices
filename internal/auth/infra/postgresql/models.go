// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package postgresql

type User struct {
	ID       int32
	Name     string
	Email    string
	Password string
}
