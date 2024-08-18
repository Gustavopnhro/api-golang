package main

import "github.com/Gustavopnhro/api-golang/configs"

func main() {
	config, _ := configs.LoadConfig()
	println(config.DBDriver)

}
