package main

import (
	"github.com/panwenbin/ghttplog/databases"
	"github.com/panwenbin/ghttplog/routers"
)

func main() {
	databases.AutoMigrate()

	r := routers.Load()
	r.Run()
}
