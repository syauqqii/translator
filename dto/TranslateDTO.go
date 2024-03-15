package dto

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
