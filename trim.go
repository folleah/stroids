package stroid

import "strings"

func (s *Str) Trim(cutset ...string) *Str {
	if len(cutset) == 0 {
		s.value = strings.TrimSpace(s.value)
	} else {
		s.value = strings.Trim(s.value, cutset[0])
	}

	return s
}

func (s *Str) TrimStart(cutset ...string) *Str {
	if len(cutset) == 0 {
		s.value = strings.TrimLeft(s.value, " ")
	} else {
		s.value = strings.TrimLeft(s.value, cutset[0])
	}

	return s
}

func (s *Str) TrimEnd(cutset ...string) *Str {
	if len(cutset) == 0 {
		s.value = strings.TrimRight(s.value, " ")
	} else {
		s.value = strings.TrimRight(s.value, cutset[0])
	}

	return s
}
