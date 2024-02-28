package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	letsgo "github.com/bas24/googletranslatefree"
)

// TranslateRequest => menampung request dari user
type TranslateRequest struct {
	Text           string `json:"text"`
	TargetLanguage string `json:"target_language"`
	SourceLanguage string `json:"source_language"`
}

// TranslateResponse => menampung hasil translate
type TranslateResponse struct {
	OriginalText   string `json:"original_text"`
	TranslatedText string `json:"translated_text"`
	SourceLanguage string `json:"source_language"`
	TargetLanguage string `json:"target_language"`
}

func translateHandler(w http.ResponseWriter, r *http.Request) {
	// Formating JSON request dari user
	var request TranslateRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Proses melakukan translate
	translation, err := letsgo.Translate(
		request.Text,
		request.SourceLanguage,
		request.TargetLanguage,
	)
	if err != nil {
		http.Error(w, fmt.Sprintf("Translation failed: %v", err), http.StatusInternalServerError)
		return
	}

	// Memasukkan hasil translate ke reseponse translate
	response := TranslateResponse{
		OriginalText:   request.Text,
		TranslatedText: translation,
		SourceLanguage: request.SourceLanguage,
		TargetLanguage: request.TargetLanguage,
	}

	// Lempar hasil transalte
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/translate", translateHandler)

	port := 6666
	fmt.Printf("Server is running at http://localhost:%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
