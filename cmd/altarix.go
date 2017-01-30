package main

import (
	"altarix_test/database"
	"altarix_test/web"
	"altarix_test/socket"
)

func main() {
	database.InitDB()
	defer database.DataSource.Close()
	go socket.Run()
	web.Run()
}
