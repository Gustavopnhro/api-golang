package main

import (
	"github.com/Gustavopnhro/api-golang/configs"
)

func main() {
	cfg, _ := configs.LoadConfig(".")
	println(cfg.DBDriver)

}
