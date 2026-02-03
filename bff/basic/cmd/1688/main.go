package main

import (
	"fmt"

	"gitee.com/zuyanlongnb666/crawl/Upload"
	"gitee.com/zuyanlongnb666/crawl/cra"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func main() {
	Mysql()
	crawl := cra.Crawl(
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36",
		"https://sale.1688.com/factory/u0vjcc4j.html?spm=a260k.home2025.centralDoor.ddoor.481e3597kOvdkT&topOfferIds=1002101950199")
	for _, m := range crawl {
		crt := Crt{
			Title: m["title"],
			Img:   m["src"],
			Price: m["price"],
		}
		Upload.Upload(m["src"], "222")
		err = crt.CreateAdd(DB)
		if err != nil {
			return
		}
		fmt.Println("添加成功")
	}
}
func Mysql() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:4ay1nkal3u8ed77y@tcp(115.190.43.83:3306)/p2308a?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println("数据库连接成功")
	err = DB.AutoMigrate(Crt{})
	if err != nil {
		return
	}
	fmt.Println("表迁移成功")

}

type Crt struct {
	Title string `grom:"type:varchar(30)"`
	Img   string `grom:"type:varchar(50)"`
	Price string `grom:"type:varchar(30)"`
}

func (c *Crt) CreateAdd(db *gorm.DB) error {
	return db.Debug().Create(&c).Error
}
