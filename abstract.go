package stroids

const (
	LF   = "\n"
	CRLF = "\r\n"
)

type StrUnicode struct {
	value string
}

type StrEntry struct {
	Str          string
	IncludeEntry bool
}

type StrPos int

type Offset struct {
	From StrPos
	To   StrPos
}

type StrTruncate struct {
}

func (s *StrUnicode) String() string {
	return s.value
}

func (s *StrUnicode) Length() int {
	return len(s.value)
}

func (s *StrUnicode) ToRunes() []rune {
	return []rune(s.value)
}

func (s *StrUnicode) Chunk(length int) []string {
	chars := s.ToRunes()
	var chunked []string
	var from int

	for i := range chars {
		if i%length == 0 {
			chunked = append(chunked, string(chars[from:i]))
			from = i
		}
	}

	chunked = append(chunked, string(chars[from:]))

	return chunked
}

func (strPos *StrPos) handy(source string) (res StrPos) {
	typedLen := StrPos(len(source))
	if *strPos < 0 {
		res = typedLen + *strPos

		if res < 0 {
			res = 0
		}

		return res
	}

	if *strPos > typedLen {
		return typedLen
	}

	return *strPos
}

func (offset Offset) handy(source string) Offset {
	offset.From = offset.From.handy(source)
	offset.To = offset.To.handy(source)

	return offset
}
