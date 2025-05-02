package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(text string) string {
	if isMorse(text) {
		return morse.ToText(text)
	}
	return morse.ToMorse(text)
}

func isMorse(text string) bool {
	text = strings.ReplaceAll(text, " ", "")
	morseChars := ".-"

	for _, char := range text {
		if !strings.ContainsRune(morseChars, char) {
			return false
		}
	}
	return true
}
