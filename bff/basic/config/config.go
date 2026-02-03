package config

import (
	__ "yuekao/srv/basic/proto"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	UserClient __.UserClient
	DB         *gorm.DB
	Rdb        *redis.Client
)
