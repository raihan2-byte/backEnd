package barang

type InputBarang struct {
	Nama         string `form:"nama" binding:"required"`
	Harga        int    `form:"harga" binding:"required"`
	CategoryID   int    `form:"categoryID" binding:"required"`
	CategoryData Category
}

type GetBarang struct {
	ID int `uri:"id" binding:"required"`
}

type GetCategory struct {
	Category int `uri:"id" binding:"required"`
}