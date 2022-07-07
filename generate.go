package stroid

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func FromString(str string) Str {
	return Str{
		value: str,
	}
}

func FromRand(length int, characters ...string) Str {
	var chars string
	if len(characters) > 0 {
		chars = characters[0]
	} else {
		chars = charset
	}

	charactersRunes := []rune(chars)
	charactersRunesLen := len(charactersRunes)

	seedRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	randRunes := make([]rune, length)
	for i := range randRunes {
		randRunes[i] = charactersRunes[seedRand.Intn(charactersRunesLen)]
	}

	return Str{
		value: string(randRunes),
	}
}