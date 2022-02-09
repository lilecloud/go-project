package main

import (
	"fmt"
	"./cache"
)

func main() {
	cache := cache.NewCache()

	cache.Put("lile", "test")
	fmt.Println(cache.Get("lile"))
}
