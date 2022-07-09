package stroids

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"math"
	"regexp"
	"strings"
)

type UnicodeStr struct {
	value string
}

type Offset int

type Needle struct {
	Str           string
	IncludeNeedle bool
}

type Width struct {
	From Offset
	To   Offset
}

type IndexValue struct {
	Str    string
	Offset Offset
}

func Init() {
	fmt.Println(cases.Caser{})
	fmt.Println(language.Afrikaans)
}

func (s *UnicodeStr) String() string {
	return s.value
}

func (s *UnicodeStr) StartsWith(str string) bool {
	return strings.HasPrefix(s.value, str)
}

func (s *UnicodeStr) EndsWith(str string) bool {
	return strings.HasSuffix(s.value, str)
}

func (s *UnicodeStr) EqualsTo(str string) bool {
	return strings.Compare(s.value, str) == 0
}

func (s *UnicodeStr) Match(pattern string) (bool, error) {
	return regexp.MatchString(pattern, s.value)
}

func (s *UnicodeStr) ContainsAny(strs []string) bool {
	for _, str := range strs {
		if strings.Contains(s.value, str) {
			return true
		}
	}

	return false
}

func (s UnicodeStr) IndexOf(indexValue IndexValue) int {
	absOffset := int(math.Abs(float64(indexValue.Offset)))
	if absOffset >= s.Length() || absOffset == 0 {
		return strings.Index(s.value, indexValue.Str)
	}

	//s.Split()
	//for _, char := range []rune(s.value) {
	//	char
	//}

	return 1
}

func (s *UnicodeStr) Length() int {
	return len(s.value)
}

func (s UnicodeStr) ToRunes() []rune {
	return []rune(s.value)
}

func (o Offset) handy(source string) (res int) {
	typedLen := Offset(len(source))
	if o < 0 {
		res = int(typedLen + o)

		if res < 0 {
			res = 0
		}

		return res
	}

	if o > typedLen {
		return int(typedLen)
	}

	return int(o)
}
