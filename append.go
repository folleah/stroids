package stroid

func (s *Str) Append(str string) *Str {
	s.value += str

	return s
}

func (s *Str) Prepend(str string) *Str {
	s.value = str + s.value

	return s
}

func (s *Str) EnsureStart(str string, removeDoubles bool) *Str {
	if !s.StartsWith(str) {
		return s.Prepend(str)
	} else if removeDoubles {
		// todo реализовать функционал по удалению дублей в начале
	}

	return s
}

func (s *Str) EnsureEnd(str string, removeDoubles bool) *Str {
	if !s.EndsWith(str) {
		return s.Append(str)
	} else if removeDoubles {
		// todo реализовать функционал по удалению дублей в конце
	}

	return s
}

func (s *Str) Repeat(times int) *Str {
	for i := 1; i < times; i++ {
		s.Append(s.value)
	}

	return s
}
