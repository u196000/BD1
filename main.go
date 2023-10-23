package main

import (
	"gorillacosa/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
    r := mux.NewRouter()

    // Asignar el manejador GetUsuariosHandler a la ruta /usuarios
    r.HandleFunc("/usuarios", handler.GetUsuariosHandler).Methods("GET")

    http.Handle("/", r)

    // Iniciar el servidor
    http.ListenAndServe(":8080", nil)
}