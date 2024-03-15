package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"translator/dto"
	"translator/service"
)

// TranslateHandler => menangani permintaan translate
func TranslateHandler(w http.ResponseWriter, r *http.Request) {
	var request dto.TranslateRequest

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, " ! ERROR: Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Validate request fields
	if  request.Text == "" || request.TargetLanguage == "" || request.SourceLanguage == ""{
		http.Error(w, " ! ERROR: Text and TargetLanguage are required", http.StatusBadRequest)
		return
	}

	// Process translation
	translation, err := service.Translate(request)
	if err != nil {
		http.Error(w, fmt.Sprintf(" ! ERROR: Translation failed: %v", err), http.StatusInternalServerError)
		return
	}

	// Prepare response
	response := dto.TranslateResponse{
		OriginalText:   request.Text,
		TranslatedText: translation,
		SourceLanguage: request.SourceLanguage,
		TargetLanguage: request.TargetLanguage,
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf(" ! ERROR: Error encoding JSON: %v", err), http.StatusInternalServerError)
		return
	}
}
