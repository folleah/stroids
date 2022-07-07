package stroid

import "strings"

func (s *Str) Before(needle Needle) *Str {
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

func (s *Str) BeforeLast(needle Needle) *Str {
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

func (s *Str) After(needle Needle) *Str {
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

func (s *Str) AfterLast(needle Needle) *Str {
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

func (s *Str) Slice(size Width) *Str {
	from := size.From.handy(s.value)
	to := size.To.handy(s.value)

	if from >= to {
		s.value = ""
	} else {
		s.value = string(s.ToRunes()[from:to])
	}

	return s
}
