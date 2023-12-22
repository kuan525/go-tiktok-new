package dal

import (
	"go-tiktok-new/biz/dal/db"
	"go-tiktok-new/biz/mw/redis"
)

func Init() {
	db.Init()
	redis.InitRedis()
}
