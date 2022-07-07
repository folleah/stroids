package stroid

func (s *Str) PadBoth(width int, char rune) *Str {
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

func (s *Str) PadStart(width int, char rune) *Str {
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

func (s *Str) PadEnd(width int, char rune) *Str {
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
