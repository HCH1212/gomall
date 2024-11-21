package dal

import (
	"github.com/HCH1212/gomall/app/frontend/biz/dal/mysql"
	"github.com/HCH1212/gomall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
