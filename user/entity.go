package user

import "time"

type User struct {
	ID        int
	Username  string
	Email     string
	Password  string
	Role int
	CreatedAt time.Time
	UpdatedAt time.Time
}
