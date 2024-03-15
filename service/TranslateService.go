package service

import (
	"translator/dto"

	letsgo "github.com/bas24/googletranslatefree"
)

// Translate => melakukan translate teks
func Translate(request dto.TranslateRequest) (string, error) {
	return letsgo.Translate(request.Text, request.SourceLanguage, request.TargetLanguage)
}
