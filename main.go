package main

import (
	"DemoApiPostgre/config"
	"DemoApiPostgre/database"
	"DemoApiPostgre/routes"
	"fmt"
)

func main() {
	config.Init()
	db, err := database.DBConnection(config.Host, config.User, config.Password, config.Port, config.DBName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	routes.Route(db)
}
