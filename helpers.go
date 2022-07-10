package stroids

import (
	"strings"
)

type Wrap struct {
	Width      int
	Delimiter  string
	WrapLetter bool
}

func (s StrUnicode) Trim(cutset ...string) StrUnicode {
	if len(cutset) == 0 {
		s.value = strings.TrimSpace(s.value)
	} else {
		s.value = strings.Trim(s.value, cutset[0])
	}

	return s
}

func (s StrUnicode) TrimStart(cutset ...string) StrUnicode {
	if len(cutset) == 0 {
		s.value = strings.TrimLeft(s.value, " ")
	} else {
		s.value = strings.TrimLeft(s.value, cutset[0])
	}

	return s
}

func (s StrUnicode) TrimEnd(cutset ...string) StrUnicode {
	if len(cutset) == 0 {
		s.value = strings.TrimRight(s.value, " ")
	} else {
		s.value = strings.TrimRight(s.value, cutset[0])
	}

	return s
}

func (s StrUnicode) PadBoth(width int, char rune) StrUnicode {
	if s.Length() >= width {
		return s
	}

	for i := 0; s.Length() < width; i++ {
		if i == 0 {
			s.Append(" ")
		} else if i == 1 {
			s.Prepend(" ")
		} else if i%2 == 0 {
			s.Append(string(char))
		} else {
			s.Prepend(string(char))
		}
	}

	return s
}

func (s StrUnicode) PadStart(width int, char rune) StrUnicode {
	if s.Length() >= width {
		return s
	}

	for i := 0; s.Length() < width; i++ {
		if i == 0 {
			s.Prepend(" ")
		} else {
			s.Prepend(string(char))
		}
	}

	return s
}

func (s StrUnicode) PadEnd(width int, char rune) StrUnicode {
	if s.Length() >= width {
		return s
	}

	for i := 0; s.Length() < width; i++ {
		if i == 0 {
			s.Append(" ")
		} else {
			s.Append(string(char))
		}
	}

	return s
}

func (s StrUnicode) WordWrap(wrap Wrap) StrUnicode {
	var formatted []rune

	for i, char := range s.ToRunes() {
		formatted = append(formatted, char)
		if i > 0 && i%wrap.Width == 0 {
			for _, del := range []rune(wrap.Delimiter) {
				formatted = append(formatted, del)
			}
		}
	}

	s.value = string(formatted)

	return s
}

func (s StrUnicode) Mask(symbol rune, offset Offset) StrUnicode {
	offset = offset.handy(s.value)

	chars := s.ToRunes()

	var formatted []rune
	for i, char := range chars {
		if i < int(offset.From) || i > int(offset.To) {
			formatted = append(formatted, char)
		} else {
			formatted = append(formatted, symbol)
		}
	}

	s.value = string(formatted)

	return s
}

func (s StrUnicode) Reverse() StrUnicode {
	runes := s.ToRunes()
	if len(runes) == 0 {
		return s
	}

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	s.value = string(runes)

	return s
}
