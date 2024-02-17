package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	// local docker 用
	redisHost := "localhost:6379"
	redisPassword := ""
	op := &redis.Options{Addr: redisHost, Password: redisPassword}

	// azure 用
	// password = アクセスキー
	// https://learn.microsoft.com/ja-jp/azure/azure-cache-for-redis/cache-go-get-started
	// redisHost := os.Getenv("REDIS_HOST")
	// redisHost = redisHost + ":6380"
	// redisPassword := os.Getenv("REDIS_PASSWORD")
	// op := &redis.Options{Addr: redisHost, Password: redisPassword, TLSConfig: &tls.Config{MinVersion: tls.VersionTLS12}}

	fmt.Println("ホスト名：", redisHost)
	fmt.Println("パスワード：", redisPassword)

	// Redis クライアントを作成
	rdb := redis.NewClient(op)

	// Ping で接続を確認
	err := rdb.Ping(ctx).Err()
	if err != nil {
		log.Fatalf("failed to connect with redis instance at %s - %v", redisHost, err)
	}

	// DB にデータをセットし、取得する
	err = rdb.Set(ctx, "key1", "value111", 0).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("key1 data is set")

	val, err := rdb.Get(ctx, "key1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key1 value is", val)

	// DBにないkeyを取得しようとするとエラーが出る
	// val2, err := rdb.Get(ctx, "key2").Result()
	// if err == redis.Nil {
	// 	fmt.Println("key2 does not exist")
	// } else if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("key2", val2)
	// }
	// // Output: key value
	// // key2 does not exist

	// DBの中身を全て表示
	keys, err := rdb.Keys(ctx, "*").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("All the data in the database0 are below:")
	for _, key := range keys {
		fmt.Println(key)
	}
	fmt.Println("Finished")

}
