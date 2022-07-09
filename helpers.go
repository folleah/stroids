package stroids

import "strings"

func (s *UnicodeStr) Trim(cutset ...string) *UnicodeStr {
	if len(cutset) == 0 {
		s.value = strings.TrimSpace(s.value)
	} else {
		s.value = strings.Trim(s.value, cutset[0])
	}

	return s
}

func (s *UnicodeStr) TrimStart(cutset ...string) *UnicodeStr {
	if len(cutset) == 0 {
		s.value = strings.TrimLeft(s.value, " ")
	} else {
		s.value = strings.TrimLeft(s.value, cutset[0])
	}

	return s
}

func (s *UnicodeStr) TrimEnd(cutset ...string) *UnicodeStr {
	if len(cutset) == 0 {
		s.value = strings.TrimRight(s.value, " ")
	} else {
		s.value = strings.TrimRight(s.value, cutset[0])
	}

	return s
}

func (s *UnicodeStr) AddSpaces() *UnicodeStr {
	// todo добавить пробелы после всех знаков, если их нет. Учесть конец строки

	return s
}

func (s *UnicodeStr) PadBoth(width int, char rune) *UnicodeStr {
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

func (s *UnicodeStr) PadStart(width int, char rune) *UnicodeStr {
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

func (s *UnicodeStr) PadEnd(width int, char rune) *UnicodeStr {
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
