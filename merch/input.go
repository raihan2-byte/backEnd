package merch

type CreateMerch struct {
	Name  string `form:"name" binding:"required"`
	Price int    `form:"price" binding:"required"`
}

type GetMerch struct {
	ID int `uri:"id" binding:"required"`
}
