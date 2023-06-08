package util

import (
	"encoding/json"
	"log"

	redis "gopkg.in/redis.v5"
)

var client *redis.Client

func InitialiazeRedis() {
	log.Println("LOG SAMERE REDIS")
	client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	SetBookList(BookCollection)
}

func SetBookList(bookList *Collection) {
	if err := client.Set("booklist", bookList, 0).Err(); err != nil {
		log.Println("Couldn't save the booklist, Err:", err)
	}
}

func GetBookList() *Collection {
	val, err := client.Get("booklist").Result()
	if err != nil {
		panic(err)
	}
	var books Collection
	if err = json.Unmarshal([]byte(val), &books); err != nil {
		panic(err)
	}
	return &books
}
