package dal

import (
	"github.com/HCH1212/gomall/app/user/biz/dal/mysql"
	"github.com/HCH1212/gomall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
