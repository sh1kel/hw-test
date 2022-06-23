package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")
var ErrFirstSpace = errors.New("first space")

func Unpack(packedString string) (string, error) {
	outString := strings.Builder{}
	var currentSymbol rune
	var previousSymbol rune
	var repeat int
	var strLen = len([]rune(packedString))
	var err error

	for i, r := range packedString {
		// fmt.Printf("%d Processing %c\n", i, r)
		if i == 0 && unicode.IsDigit(r) { // Первый символ - цифра, выходим
			return "", ErrInvalidString
		}
		if i == 0 && r == ' ' { // Первый символ - пробел, выходим
			return "", ErrFirstSpace
		}
		if i == 0 && !unicode.IsDigit(r) { // Первый символ - не цифра
			currentSymbol = r
		}
		if i > 0 {
			previousSymbol = currentSymbol
			currentSymbol = r
			switch {
			case unicode.IsDigit(currentSymbol) && unicode.IsDigit(previousSymbol): // две цифры подряд
				return "", ErrInvalidString
			case unicode.IsDigit(currentSymbol): // && previousSymbol != '\\': // Текущий - цифра, предыдущий - не слэш
				repeat, err = strconv.Atoi(string(r))
				if err != nil {
					return "", err
				}
				// fmt.Printf("Repeat %c %d times\n", previousSymbol, repeat)
				for i := 1; i <= repeat; i++ {
					// fmt.Printf("%d Save %c\n", i, previousSymbol)
					outString.WriteRune(previousSymbol)
				}
			case !unicode.IsDigit(currentSymbol) && !unicode.IsDigit(previousSymbol):
				// текущий не цифра, предыдыущий не цифра
				// fmt.Printf("%d Save %c\n", i, previousSymbol)
				outString.WriteRune(previousSymbol)
				if i == strLen-1 {
					// fmt.Printf("%d Save %c\n", i, currentSymbol)
					outString.WriteRune(currentSymbol)
				}
			case !unicode.IsDigit(currentSymbol) && unicode.IsDigit(previousSymbol):
				// текущий не цифра, предыдущий - цифра
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
