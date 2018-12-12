package main

import (
	"goRest/app/config"
	"log"
)

func main() {
	c := new(config.Config)
	err := config.LoadConfigFromJSON(c)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(c.Get("DB.DATABASE"), c.Get("DB.PORT"))
	// sql := new(database.MySQL)
}
