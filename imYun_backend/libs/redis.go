package libs

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
)

var Ctx = context.Background()
var Rdb0 *redis.Client
var Rdb1 *redis.Client
var Rdb2 *redis.Client


func CreateRedisDB()  {
	if GoEnv:=os.Getenv("GORUNENV");GoEnv=="docker"{
		Rdb0 = redis.NewClient(&redis.Options{
			Addr	:	"redis:6379",
			DB		:	0,
		})
		Rdb1 = redis.NewClient(&redis.Options{
			Addr	:	"redis:6379",
			DB		:	1,
		})
		Rdb2 = redis.NewClient(&redis.Options{
			Addr	:	"redis:6379",
			DB		:	2,
		})
	}else{
		var c YConfig
		data := c.GetConfig()
		Rdb0 = redis.NewClient(&redis.Options{
			Addr	:	data.Redis.Host+":"+data.Redis.Port,
			Password:   data.Redis.PassWord,
			DB		:	data.Redis.DB0,
		})
		Rdb1 = redis.NewClient(&redis.Options{
			Addr	:	data.Redis.Host+":"+data.Redis.Port,
			Password:   data.Redis.PassWord,
			DB		:	data.Redis.DB1,
		})
		Rdb2 = redis.NewClient(&redis.Options{
			Addr	:	data.Redis.Host+":"+data.Redis.Port,
			Password:   data.Redis.PassWord,
			DB		:	data.Redis.DB2,
		})
	}
	pong , err := Rdb0.Ping(Ctx).Result()
	if err != nil{
		fmt.Println("redis 0 连接失败", pong, err)
	}

	pong , err = Rdb1.Ping(Ctx).Result()
	if err != nil{
		fmt.Println("redis 1 连接失败", pong, err)
	}

	pong , err = Rdb2.Ping(Ctx).Result()
	if err != nil{
		fmt.Println("redis 2 连接失败", pong, err)
	}
	fmt.Println("redis 连接成功", pong)
}

func GetRedisDB0() *redis.Client {
	return Rdb0
}
func GetRedisDB1() *redis.Client {
	return Rdb1
}
func GetRedisDB2() *redis.Client {
	return Rdb2
}