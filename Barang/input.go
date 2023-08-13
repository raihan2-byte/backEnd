package barang

type InputBarang struct {
	NamaPaket    *string `form:"nama_paket"`
	Nama         string  `form:"nama" binding:"required"`
	HargaAwal    *int    `form:"harga_awal"`
	Harga        int     `form:"harga" binding:"required"`
	CategoryID   int     `form:"categoryID" binding:"required"`
	CategoryData Category
}

type GetBarang struct {
	ID int `uri:"id" binding:"required"`
}

type GetCategory struct {
	Category int `uri:"id" binding:"required"`
}