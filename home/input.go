package home

type CreateTagHome struct {
	Heading string `json:"heading" binding:"required"`
	Text    string `json:"text" binding:"required"`
}

type GetIdTagHome struct {
	ID int `uri:"id" bindind:"required"`
}