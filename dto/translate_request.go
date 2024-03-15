package dto

// TranslateRequest => menampung request dari user
type TranslateRequest struct {
	Text           string `json:"text"`
	TargetLanguage string `json:"target_language"`
	SourceLanguage string `json:"source_language"`
}
