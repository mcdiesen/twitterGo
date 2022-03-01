package main

import (
	"fmt"
	_ "github.com/mcdiesen/twitterGo/bd"
	_ "github.com/mcdiesen/twitterGo/handler"
	"log"
	"twitterGo/bd"
	"twitterGo/handlers"

	//"twitterGo/bd"
	//"twitterGo/handlers"
)

func main() {
	if bd.ChequeoConection()==0 {
		log.Fatal("Sin conexión a la base de datos")
		return
	}
	handlers.Manejadores()
}
