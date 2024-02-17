package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "key1", "value111", 0).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("key1 data is set")

	val, err := rdb.Get(ctx, "key1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key1 value is", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
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
