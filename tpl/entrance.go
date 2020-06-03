package tpl

var Entrance = `package main

import (
	"log"
	"<module>/config"
	"<module>/storage"
	"<module>/router"
)

func main() {
	// configure initialize
	config.Init()

	// DB init
	db, err := storage.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// router init
	router.Init()
}
`
