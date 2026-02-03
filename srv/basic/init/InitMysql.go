package init

import (
	"fmt"
	"yuekao/srv/basic/config"
	"yuekao/srv/handler/model"

	"github.com/gospacex/gospacex/core/storage/conf"
	"github.com/gospacex/gospacex/core/storage/db/mysql"
)

var err error

func InitMysql() {
	config.DB, err = mysql.Init(true, "debug", conf.Cfg.Mysql)
	if err != nil {
		return
	}
	fmt.Println("数据库连接成功")
	err = config.DB.AutoMigrate(&model.User{}, &model.Goods{})
	if err != nil {
		return
	}
	fmt.Println("表迁移成功")
}
