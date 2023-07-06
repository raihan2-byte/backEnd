package barang

type InputBarang struct {
	Nama  string `form:"nama" binding:"required"`
	Harga int    `form:"harga" binding:"required"`
}

type GetBarang struct {
	ID int `uri:"id" binding:"required"`
}