package stroids

import (
	"regexp"
	"strings"
)

func (s *StrUnicode) StartsWith(str string) bool {
	return strings.HasPrefix(s.value, str)
}

func (s *StrUnicode) EndsWith(str string) bool {
	return strings.HasSuffix(s.value, str)
}

func (s *StrUnicode) EqualsTo(str string) bool {
	return strings.Compare(s.value, str) == 0
}

func (s *StrUnicode) Match(pattern string) (bool, error) {
	return regexp.MatchString(pattern, s.value)
}

func (s *StrUnicode) ContainsAny(strs []string) bool {
	for _, str := range strs {
		if strings.Contains(s.value, str) {
			return true
		}
	}

	return false
}

func (s *StrUnicode) Index(str string) int {
	return strings.Index(s.value, str)
}

func (s *StrUnicode) LastIndex(str string) int {
	return strings.LastIndex(s.value, str)
}
