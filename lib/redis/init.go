package redis

import (
	"time"

	"github.com/astaxie/beego"
)

var Client *RedisManager

func Init(key string) (err error) {
	host := beego.AppConfig.String("redis::host")
	auth := beego.AppConfig.String("redis::auth")
	poolsize, _ := beego.AppConfig.Int("redis::poolsize")
	timeout, _ := beego.AppConfig.Int("redis::timeout")
	if len(key) > 0 && len(auth) > 0 {
		if err != nil {
			return
		}
	}
	Client, err = NewRedisManager(
		host, auth, poolsize, time.Millisecond*time.Duration(timeout))
	return
}
