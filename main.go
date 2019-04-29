package main

import (
	"github.com/go-redis/redis"
	"fmt"
	"crypto/tls"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "", // no password set
		DB:       2,  // use default DB
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	})

	clientLocal := redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "", // no password set
		DB:       2,  // use default DB
	})
	
	
	keys, err := client.Keys("*").Result()
	if err != nil {
		panic(err)
	}
	for _, key := range keys {
		hashKeys, err :=client.HKeys(key).Result()

		if err != nil {
			panic(err)
		}

		for _, hashKey := range hashKeys {
			value, err := client.HGet(key, hashKey).Result()

			if err != nil {
				panic(err)
			}

			result, err := clientLocal.HSet(key, hashKey, value).Result()
			if err != nil {
				panic(err)
			}
			fmt.Printf("Saving the key %s with hash %s -%t\n", key, hashKey, result)
		}
	}
}