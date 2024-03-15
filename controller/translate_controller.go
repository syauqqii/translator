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

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, " ! ERROR: Invalid JSON format", http.StatusBadRequest)
		return
	}

	translation, err := service.Translate(request)
	if err != nil {
		http.Error(w, fmt.Sprintf(" ! ERROR: Translation failed: %v", err), http.StatusInternalServerError)
		return
	}

	response := dto.TranslateResponse{
		OriginalText:   request.Text,
		TranslatedText: translation,
		SourceLanguage: request.SourceLanguage,
		TargetLanguage: request.TargetLanguage,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf(" ! ERROR: Error encoding JSON: %v", err), http.StatusInternalServerError)
		return
	}
}
