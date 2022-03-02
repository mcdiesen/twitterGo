package handlers

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"twitterGo/middlew"
	"twitterGo/routers"
)

/*Manejadores seteo mi puerto el handler y pongo a escuchar al server maneja las rutas*/
func Manejadores() {
	router := mux.NewRouter()

	//registro de usuarios
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

	//rutas que menejaremos
}
