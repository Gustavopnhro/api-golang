package main

import (
	"github.com/Gustavopnhro/api-golang/internal/database"
)

func main() {
	// config, err := configs.LoadConfig()
	// if err != nil {
	// log.Fatal("cannot load config:", err)
	// }

	// println(config.DBDriver)
	database.InitDB()
}
