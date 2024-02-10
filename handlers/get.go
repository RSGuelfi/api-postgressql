package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/RSGuelfi/models"
)

func Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(char.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	todo, err := models.Get(int64(id))
	if err != nil {
		log.Printf("Erro ao obter registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Context-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}
