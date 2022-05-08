// package main

// import (
// 	"fmt"

// 	_ "github.com/go-redis"
// 	_ "github.com/go-redis/redis"
// 	//"github.com/go-redis/redis@v6.15.9+incompatible"
// 	//_ "github.com/go-redis/redis@v6.15.9+incompatible"
// )

// var redisdb *redis.Client

// // 初始化连接
// func initRedis() (err error) {
// 	redisdb = redis.NewClient(&redis.Options{
// 		Addr:     "127.0.0.1:6379",
// 		Password: "", // no password set
// 		DB:       0,  // use default DB
// 	})

// 	_, err = redisdb.Ping().Result()
// 	return

// }

// func main() {
// 	err := initRedis()
// 	if err != nil {
// 		fmt.Printf("connect redis failed, err: %v\n", err)
// 		return
// 	}
// 	fmt.Println("连接成功！")
// }

package main

import "fmt"

func main() {
	fmt.Println("hello")
}
