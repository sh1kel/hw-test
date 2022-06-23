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
	var repeat int
	strLen := len([]rune(packedString))
	var err error

	for i, r := range packedString {
		// fmt.Printf("%d Processing %c\n", i, r)
		if i == 0 {
			switch {
			case unicode.IsDigit(r):
				return "", ErrInvalidString
			case r == ' ':
				return "", ErrFirstSpace
			case unicode.IsLetter(r):
				currentSymbol = r
			}
		}

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
				// fmt.Printf("Repeat %c %d times\n", previousSymbol, repeat)
				for i := 1; i <= repeat; i++ {
					// fmt.Printf("%d Save %c\n", i, previousSymbol)
					outString.WriteRune(previousSymbol)
				}
			case unicode.IsLetter(currentSymbol) && unicode.IsLetter(previousSymbol):
				// текущий буква, предыдыущий буква
				// fmt.Printf("%d Save %c\n", i, previousSymbol)
				outString.WriteRune(previousSymbol)
				if i == strLen-1 {
					// fmt.Printf("%d Save %c\n", i, currentSymbol)
					outString.WriteRune(currentSymbol)
				}
			case unicode.IsLetter(currentSymbol) && unicode.IsDigit(previousSymbol):
				// текущий буква, предыдущий - цифра
				if i == strLen-1 {
					// fmt.Printf("%d Save %c\n", i, currentSymbol)
					outString.WriteRune(currentSymbol)
				}
			}
		}
	}
	// fmt.Printf("%s => %s\n", packedString, outString.String())
	return outString.String(), nil
}
