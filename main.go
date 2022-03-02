package main

import (
	//"fmt"
	//"twitterGo/handlers"

	//"github.com/mcdiesen/twitterGo/bd"
	//_ "github.com/mcdiesen/twitterGo/handlers"
	"log"
	"twitterGo/bd"
	"twitterGo/handlers"
	//"twitterGo/bd"
	//"twitterGo/handlers"
)

func main() {
	if bd.ChequeoConection() == 0 {
		log.Fatal("Sin conexión a la base de datos")
		return
	}
	handlers.Manejadores()
}
