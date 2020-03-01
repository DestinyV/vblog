package main

import (
	"vblog.me.com/GinServer/Router"

	Mysql "vblog.me.com/GinServer/Databases"
)

func main() {
	defer Mysql.DB.Close()
	Router.InitRouter()
}
