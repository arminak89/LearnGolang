package main

import (
	"log"
	"sync"
)

type Config struct {
	Name string
}

var singletonConfig *Config
var doOnce sync.Once

// GetConfig returns a base Config
func GetConfig() *Config {
	doOnce.Do(func() {
		log.Printf("initizlizing config struct")
		singletonConfig = &Config{
			Name: "default_name",
		}
	})
	return singletonConfig
}
func main() {
	for i := 1; i < 5; i++ {
		log.Printf("[%d] calling GetConfig()", i)
		GetConfig()
	}
}

// OUTPUT:
// 2021/01/02 11:45:02 [1] calling GetConfig()
// 2021/01/02 11:45:02 initizlizing config struct
// 2021/01/02 11:45:02 [2] calling GetConfig()
// 2021/01/02 11:45:02 [3] calling GetConfig()
// 2021/01/02 11:45:02 [4] calling GetConfig()

// Despite calling GetConfig multiple times, initialization only happened once.
