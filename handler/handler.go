package handler

import (
	"encoding/json"
	"gorillacosa/models"
	"net/http"
)

func GetUsuariosHandler(w http.ResponseWriter, r *http.Request) {
    // Convertir el slice de usuarios en JSON
    jsonUsuarios, err := json.Marshal(models.Servicios)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Establecer la cabecera Content-Type
    w.Header().Set("Content-Type", "application/json")

    // Escribir el JSON en la respuesta
    w.Write(jsonUsuarios)
}