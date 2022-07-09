package stroids

import "strings"

func (s *UnicodeStr) Before(needle Needle) *UnicodeStr {
	firstIndex := strings.Index(s.value, needle.Str)
	if needle.IncludeNeedle {
		firstIndex += len(needle.Str)
	}

	if firstIndex != -1 {
		s.value = string([]rune(s.value)[0:firstIndex])
	} else {
		s.value = ""
	}

	return s
}

func (s *UnicodeStr) BeforeLast(needle Needle) *UnicodeStr {
	lastIndex := strings.LastIndex(s.value, needle.Str)
	if needle.IncludeNeedle {
		lastIndex += len(needle.Str)
	}

	if lastIndex != -1 {
		s.value = string([]rune(s.value)[0:lastIndex])
	} else {
		s.value = ""
	}

	return s
}

func (s *UnicodeStr) After(needle Needle) *UnicodeStr {
	firstIndex := strings.Index(s.value, needle.Str)
	if needle.IncludeNeedle {
		firstIndex -= len(needle.Str)
	}

	if firstIndex != -1 {
		s.value = string([]rune(s.value)[firstIndex:s.Length()])
	} else {
		s.value = ""
	}

	return s
}

func (s *UnicodeStr) AfterLast(needle Needle) *UnicodeStr {
	lastIndex := strings.LastIndex(s.value, needle.Str)
	if needle.IncludeNeedle {
		lastIndex -= len(needle.Str)
	}

	if lastIndex != -1 {
		s.value = string([]rune(s.value)[lastIndex:s.Length()])
	} else {
		s.value = ""
	}

	return s
}

func (s *UnicodeStr) Slice(size Width) *UnicodeStr {
	from := size.From.handy(s.value)
	to := size.To.handy(s.value)

	if from >= to {
		s.value = ""
	} else {
		s.value = string(s.ToRunes()[from:to])
	}

	return s
}
