package stroids

import (
	"strings"
	"unicode"
)

type Case int8

var (
	SnakeCaseLower = Case(1)
	SnakeCaseUpper = Case(2)
	KebabCase      = Case(3)
	CamelCase      = Case(10)
	CamelCaseLower = Case(11)
)

func (s *StrUnicode) Title(allWords bool) *StrUnicode {
	var prevCharacter rune

	if s.Length() == 0 {
		return s
	}

	var formatted []rune
	if allWords {
		for _, char := range s.ToRunes() {
			if !unicode.IsLetter(prevCharacter) && unicode.IsLetter(char) {
				formatted = append(formatted, unicode.ToTitle(char))
			} else {
				formatted = append(formatted, unicode.ToLower(char))
			}
			prevCharacter = char
		}
	} else {
		for i, char := range s.ToRunes() {
			if i == 0 && unicode.IsLetter(char) {
				formatted = append(formatted, unicode.ToTitle(char))
			} else if unicode.IsUpper(char) && !unicode.IsLetter(prevCharacter) || unicode.IsUpper(prevCharacter) {
				formatted = append(formatted, char)
			} else {
				formatted = append(formatted, unicode.ToLower(char))
			}
			prevCharacter = char
		}
	}

	s.value = string(formatted)

	return s
}

func (s *StrUnicode) Lower() *StrUnicode {
	s.value = strings.ToLower(s.value)

	return s
}

func (s *StrUnicode) Upper() *StrUnicode {
	s.value = strings.ToUpper(s.value)

	return s
}

func (s *StrUnicode) ToCase(c Case) *StrUnicode {
	switch c {
	case SnakeCaseLower:
		s.value = likeSnakeCase(s.value, '_', false, false)
	case SnakeCaseUpper:
		s.value = likeSnakeCase(s.value, '_', false, true)
	case KebabCase:
		s.value = likeSnakeCase(s.value, '-', true, false)
	case CamelCase:
		s.value = likeCamelCase(s.value, true)
	case CamelCaseLower:
		s.value = likeCamelCase(s.value, false)
	}

	return s
}

func likeCamelCase(str string, titleFirstLetter bool) string {
	var formatted []rune
	var prevChar rune
	for i, char := range []rune(str) {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			upperCond := unicode.IsDigit(prevChar) && unicode.IsLetter(char)
			upperCond = upperCond || unicode.IsLetter(prevChar) && unicode.IsDigit(char)
			upperCond = upperCond || !unicode.IsLetter(prevChar) && !unicode.IsDigit(prevChar)

			if !upperCond || unicode.IsLetter(char) && i == 0 && !titleFirstLetter {
				formatted = append(formatted, unicode.ToLower(char))
			} else {
				formatted = append(formatted, unicode.ToUpper(char))
			}
		}
		prevChar = char
	}
	return string(formatted)
}

func likeSnakeCase(str string, delimiter rune, titleWords, toUpper bool) string {
	var formatted []rune

	var prevChar rune
	for _, char := range []rune(str) {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			isCharUpper := false

			delimiterCond := unicode.IsDigit(prevChar) && unicode.IsLetter(char)
			delimiterCond = delimiterCond || unicode.IsLetter(prevChar) && unicode.IsLower(prevChar) && unicode.IsLetter(char) && unicode.IsUpper(char)
			delimiterCond = delimiterCond || (!unicode.IsLetter(prevChar) && !unicode.IsDigit(prevChar))
			if delimiterCond {
				if prevChar != 0 {
					formatted = append(formatted, delimiter)
				}
				isCharUpper = titleWords
			}

			if isCharUpper || toUpper {
				formatted = append(formatted, unicode.ToUpper(char))
			} else {
				formatted = append(formatted, unicode.ToLower(char))
			}
		}
		prevChar = char
	}

	return string(formatted)
}
