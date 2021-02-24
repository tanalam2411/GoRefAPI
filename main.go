package main

import (
	"ms-inventory/global"
	"ms-inventory/initialize"
)

func RunServer() {
	engine := initialize.Routers()
	global.GDB = initialize.GormMysql()
	initialize.MysqlTables(global.GDB)
	db, _ := global.GDB.DB()
	defer db.Close()

	if err := engine.Run(); err != nil {
		panic(err)
	}
}

func main() {
	RunServer()
}
