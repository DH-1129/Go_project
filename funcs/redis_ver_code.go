package funcs

import (
	"fmt"
	"time"

	"dhui.com/configs"
	"github.com/go-redis/redis"
)

// var RDB = initClient()
var Redisclient *redis.Client

func InitRedis() (err error) {
	Redisclient = redis.NewClient(&redis.Options{
		Addr:     configs.REDIS_IP + ":" + configs.MEDIA_PORT, // 指定
		Password: configs.REDIS_PASSWORD,
		DB:       configs.REDIS_DB, // redis一共16个库，指定其中一个库即可
	})
	_, err = Redisclient.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

// func initClient() *redis.Client {
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr:     "123.60.161.44:6379",
// 		Password: "123456",
// 		DB:       0,
// 	})
// 	_, err := rdb.Ping().Result()
// 	if err != nil {
// 		Danger("redis conncet fialed!")
// 		fmt.Println("eee")
// 		panic("redis conncet fialed!")
// 	}
// 	return rdb
// }

// 设置验证码过期时间
func Set_Code(key string, code string, exp time.Duration) (err error) {
	err = InitRedis()
	if err != nil {
		Danger(fmt.Sprintf("connect redis failed! err : %v\n", err))
		return err
	}
	err = Redisclient.Set(key, code, exp).Err()
	if err != nil {
		Info(fmt.Sprintf("redis set code_key err : %v\n", err))
		return err
	}
	return nil
}

// 验证验证码是否过期
func Veri_Code(code string, email string) bool {
	err := InitRedis()
	if err != nil {
		Danger(fmt.Sprintf("connect redis failed! err : %v\n", err))
		return false
	}
	res, err := Redisclient.Get(email).Result()
	if err != nil {
		Info(fmt.Sprintf("该邮箱的验证码不存在 err : %v\n", err))
		return false
	} else {
		fmt.Println(res, res == code)
		return res == code
	}
}
