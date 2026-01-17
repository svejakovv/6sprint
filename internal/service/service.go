package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func DetectAndConvert(input string) (string, error) {
	if input == "" {
		return "", errors.New("input is empty")
	}

	if isMorseCode(input) {
		result := morse.ToText(input)
		if result == "" {
			return "", errors.New("failed to decode morse code")
		}
		return result, nil
	} else {
		result := morse.ToMorse(input)
		if result == "" {
			return "", errors.New("failed to encode text to morse")
		}
		return result, nil
	}
}

func isMorseCode(s string) bool {
	trimmed := strings.NewReplacer(".", "", "-", "", " ", "", "\t", "").Replace(s)
	return trimmed == ""
}
