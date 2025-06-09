package handlers

import (
	"encoding/json"

	"go-dessert/config"
	"go-dessert/models"
	"go-dessert/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllDessertsHandler(w http.ResponseWriter, r *http.Request) {
	desserts, err := repository.GetAllDesserts(config.DB)
	if err != nil {
		http.Error(w, "Failed to retrieve desserts", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(desserts)
}

func GetDessertByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid dessert ID", http.StatusBadRequest)
		return
	}

	dessert, err := repository.GetDessertByID(config.DB, id)
	if err != nil {
		http.Error(w, "Failed to retrieve dessert", http.StatusInternalServerError)
		return
	}

	if dessert == nil {
		http.Error(w, "Dessert not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dessert)
}

func CreateDessertHandler(w http.ResponseWriter, r *http.Request) {
	var dessert models.Dessert
	err := json.NewDecoder(r.Body).Decode(&dessert)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	insertedID, err := repository.InsertDessert(config.DB, dessert)
	if err != nil {
		http.Error(w, "Failed to create dessert", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int64{"id": insertedID})
}

func UpdateDessertHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    idStr := vars["id"]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid dessert ID", http.StatusBadRequest)
        return
    }

    var dessert models.Dessert
    err = json.NewDecoder(r.Body).Decode(&dessert)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    if dessert.ID != 0 && dessert.ID != id {
        log.Printf("Warning: ID in request body (%d) does not match URL ID (%d). Using URL ID for update.", dessert.ID, id)
    }
    dessert.ID = id

    err = repository.UpdateDessert(config.DB, dessert)
    if err != nil {
        http.Error(w, "Failed to update dessert", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Dessert updated successfully"))
}

func DeleteDessertHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    idStr := vars["id"]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid dessert ID", http.StatusBadRequest)
        return
    }

    err = repository.DeleteDessert(config.DB, id)
    if err != nil {
        http.Error(w, "Failed to delete dessert", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Dessert deleted successfully"))
}