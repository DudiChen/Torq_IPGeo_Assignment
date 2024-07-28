package handlers

import (
	"Torq_IPGeo_Assignment/models"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type Handler struct {
	DB models.Store
}

func (h *Handler) GetCountry(w http.ResponseWriter, r *http.Request) {
	ip := r.URL.Query().Get("ip")
	if ip == "" {
		log.Println("Missing IP parameter")
		http.Error(w, `{"error": "missing ip parameter"}`, http.StatusBadRequest)
		return
	}

	log.Printf("Handler received loookup query for key: %v", ip)
	location, err := h.DB.FindLocation(r.Context(), ip)
	if err != nil {
		log.Printf("Failed to find location for IP %s: %v", ip, err)
		status := http.StatusInternalServerError
		if errors.Is(err, models.ErrNotFound) {
			status = http.StatusNotFound
		}

		http.Error(w, `{"error": "`+err.Error()+`"}`, status)
		return
	}

	if err := json.NewEncoder(w).Encode(location); err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, `{"error": "internal server error"}`, http.StatusInternalServerError)
	}

	locationJson, err := json.Marshal(location)
	if err != nil {
		log.Printf("Failed to convert location to json: %v", err)
	}
	log.Printf(`Store query of key "%v" responded with: %s`, ip, string(locationJson))
}
