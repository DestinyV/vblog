package main

import (
	"server/Router"

	Mysql "server/Databases"
)

func main() {
	defer Mysql.DB.Close()
	Router.InitRouter()
}
