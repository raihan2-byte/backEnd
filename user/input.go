package user

type RegisterUserInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	// Role     int    `json:"role" binding:"required"`
}

type UpdateUserInput struct {
	Username string `json:"username" binding:"required"`
	// Role     int    `json:"role" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type DeletedUser struct {
	ID int `uri:"id" binding:"required"`
}
