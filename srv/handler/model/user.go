package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `grom:"type:varchar(30)"`
	Password string `grom:"type:varchar(32)"`
}

func (u *User) FindUser(db *gorm.DB, name string) error {
	return db.Debug().Where("name = ?", name).Find(&u).Error
}

type Goods struct {
	gorm.Model
	Name  string  `grom:"type:varchar(30)"`
	Price float32 `grom:"type:decimal(10,2)"`
	Num   int     `grom:"type:int"`
}

func (g *Goods) FindGoods(db *gorm.DB, name string) error {
	return db.Debug().Where("name = ?", name).Find(&g).Error
}

func (g *Goods) GoodsAdd(db *gorm.DB) error {
	return db.Debug().Create(&g).Error
}
func (g *Goods) FindGoodsId(db *gorm.DB, id int) Goods {
	var goods Goods
	db.Debug().Where("id = ?", id).First(&goods)
	return goods
}
