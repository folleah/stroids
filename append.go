package stroids

func (s *UnicodeStr) Append(str string) *UnicodeStr {
	s.value += str

	return s
}

func (s *UnicodeStr) Prepend(str string) *UnicodeStr {
	s.value = str + s.value

	return s
}

func (s *UnicodeStr) EnsureStart(str string, removeDoubles bool) *UnicodeStr {
	if !s.StartsWith(str) {
		return s.Prepend(str)
	} else if removeDoubles {
		// todo реализовать функционал по удалению дублей в начале
	}

	return s
}

func (s *UnicodeStr) EnsureEnd(str string, removeDoubles bool) *UnicodeStr {
	if !s.EndsWith(str) {
		return s.Append(str)
	} else if removeDoubles {
		// todo реализовать функционал по удалению дублей в конце
	}

	return s
}

func (s *UnicodeStr) Repeat(times int) *UnicodeStr {
	for i := 1; i < times; i++ {
		s.Append(s.value)
	}

	return s
}
