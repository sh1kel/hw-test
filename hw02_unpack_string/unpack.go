package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var (
	ErrInvalidString = errors.New("invalid string")
	ErrFirstSpace    = errors.New("first space")
)

func Unpack(packedString string) (string, error) {
	outString := strings.Builder{}
	var currentSymbol rune
	var previousSymbol rune
	var repeat, strLen int
	var err error

	runeString := []rune(packedString)
	if strLen = len(runeString); strLen == 0 {
		return "", nil
	}

	switch {
	case unicode.IsDigit(runeString[0]):
		return "", ErrInvalidString
	case runeString[0] == ' ':
		return "", ErrFirstSpace
	case unicode.IsLetter(runeString[0]):
		currentSymbol = runeString[0]
	}

	for i, r := range packedString {
		if i > 0 {
			previousSymbol = currentSymbol
			currentSymbol = r
			switch {
			case unicode.IsDigit(currentSymbol) && unicode.IsDigit(previousSymbol): // две цифры подряд
				return "", ErrInvalidString
			case unicode.IsDigit(currentSymbol): // Текущий - цифра
				repeat, err = strconv.Atoi(string(r))
				if err != nil {
					return "", err
				}
				outString.WriteString(strings.Repeat(string(previousSymbol), repeat))
			case unicode.IsLetter(currentSymbol) && unicode.IsLetter(previousSymbol):
				// текущий буква, предыдыущий буква
				outString.WriteRune(previousSymbol)
				if i == strLen-1 {
					outString.WriteRune(currentSymbol)
				}
			case unicode.IsLetter(currentSymbol) && unicode.IsDigit(previousSymbol):
				// текущий буква, предыдущий - цифра
				if i == strLen-1 {
					outString.WriteRune(currentSymbol)
				}
			}
		}
	}
	return outString.String(), nil
}
