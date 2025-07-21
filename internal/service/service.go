package service

import (
	"regexp"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func DetermineConversionType(content string) string {
	morsePattern := regexp.MustCompile(`^[.\-\s/]+$`)
	if morsePattern.MatchString(strings.TrimSpace(content)) {
		return morse.ToText(content)
	} else {
		return morse.ToMorse(content)
	}
}