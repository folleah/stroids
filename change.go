package stroids

import (
	"regexp"
	"strings"
)

func (s StrUnicode) Append(str string) StrUnicode {
	s.value += str

	return s
}

func (s StrUnicode) Prepend(str string) StrUnicode {
	s.value = str + s.value

	return s
}

func (s StrUnicode) EnsureStart(str string, removeDoubles bool) StrUnicode {
	if !s.StartsWith(str) {
		return s.Prepend(str)
	} else if removeDoubles {
		// todo реализовать функционал по удалению дублей в начале
	}

	return s
}

func (s StrUnicode) EnsureEnd(str string, removeDoubles bool) StrUnicode {
	if !s.EndsWith(str) {
		return s.Append(str)
	} else if removeDoubles {
		// todo реализовать функционал по удалению дублей в конце
	}

	return s
}

func (s StrUnicode) Repeat(times int) StrUnicode {
	for i := 1; i < times; i++ {
		s = s.Append(s.value)
	}

	return s
}

func (s StrUnicode) Replace(old, new string, times int) StrUnicode {
	s.value = strings.Replace(s.value, old, new, times)

	return s
}

func (s StrUnicode) ReplaceAll(old, new string) StrUnicode {
	s.value = strings.ReplaceAll(s.value, old, new)

	return s
}

func (s StrUnicode) ReplaceMatch(pattern, new string) StrUnicode {
	c, _ := regexp.Compile(pattern)
	s.value = c.ReplaceAllString(s.value, new)

	return s
}

func (s StrUnicode) Before(entry StrEntry) StrUnicode {
	firstIndex := strings.Index(s.value, entry.Str)
	if entry.IncludeEntry {
		firstIndex += len(entry.Str)
	}

	if firstIndex != -1 {
		s.value = string([]rune(s.value)[0:firstIndex])
	} else {
		s.value = ""
	}

	return s
}

func (s StrUnicode) BeforeLast(entry StrEntry) StrUnicode {
	lastIndex := strings.LastIndex(s.value, entry.Str)
	if entry.IncludeEntry {
		lastIndex += len(entry.Str)
	}

	if lastIndex != -1 {
		s.value = string([]rune(s.value)[0:lastIndex])
	} else {
		s.value = ""
	}

	return s
}

func (s StrUnicode) After(entry StrEntry) StrUnicode {
	firstIndex := strings.Index(s.value, entry.Str)
	if entry.IncludeEntry {
		firstIndex -= len(entry.Str)
	}

	if firstIndex != -1 {
		s.value = string([]rune(s.value)[firstIndex:s.Length()])
	} else {
		s.value = ""
	}

	return s
}

func (s StrUnicode) AfterLast(entry StrEntry) StrUnicode {
	lastIndex := strings.LastIndex(s.value, entry.Str)
	if entry.IncludeEntry {
		lastIndex -= len(entry.Str)
	}

	if lastIndex != -1 {
		s.value = string([]rune(s.value)[lastIndex:s.Length()])
	} else {
		s.value = ""
	}

	return s
}

func (s StrUnicode) Slice(size Offset) StrUnicode {
	from := size.From.handy(s.value)
	to := size.To.handy(s.value)

	if from >= to {
		s.value = ""
	} else {
		s.value = string(s.ToRunes()[from:to])
	}

	return s
}

func (s StrUnicode) Join(strs []string) StrUnicode {
	tempVal := ""
	for i, str := range strs {
		// check that not first iteration
		if i > 0 {
			tempVal += s.value
		}

		tempVal += str
	}

	s.value = tempVal

	return s
}

func (s StrUnicode) Truncate(length int) StrUnicode {
	s.value = string(s.ToRunes()[:length])

	return s
}
