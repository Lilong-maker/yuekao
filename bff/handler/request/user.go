package request

type Login struct {
	Name     string `form:"name"  binding:"required"`
	Password string `form:"password" binding:"required"`
}
type GoodsAdd struct {
	Name  string  `form:"name"  binding:"required"`
	Price float32 `form:"price" binding:"required"`
	Num   int     `form:"num" binding:"required"`
}
type GetGoods struct {
	Id int `form:"id" binding:"required"`
}
