package stroid

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func (s *Str) Title() *Str {
	caser := cases.Title(language.English)

	s.value = caser.String(s.value)

	return s
}
