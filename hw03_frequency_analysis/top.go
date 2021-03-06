package hw03frequencyanalysis

import (
	"errors"
	"sort"
	"strings"
	"unicode"
)

type frequencyStruct struct {
	word string
	frq  int
}

var ErrFirstSpace = errors.New("empty string")

func clearString(s string) (string, error) {
	var clearS []rune
	var resultS strings.Builder
	if s == "-" {
		return "", ErrFirstSpace
	}
	clearS = []rune(s)
	for i, a := range clearS {
		if unicode.IsLetter(a) {
			resultS.WriteRune(a)
		}
		if i != len(clearS)-1 && i != 0 && a == '-' {
			resultS.WriteRune(a)
		}
	}
	return strings.ToLower(resultS.String()), nil
}

func Top10(text string) []string {
	frequencyMap := make(map[string]int)
	resultSlice := make([]string, 0, 10)
	stringSlice := strings.Fields(text)
	for _, val := range stringSlice {
		clean, err := clearString(val)
		if err != nil {
			continue
		}
		frequencyMap[clean]++
	}
	frequencySlice := make([]frequencyStruct, 0, len(frequencyMap))
	for key, val := range frequencyMap {
		frequencySlice = append(frequencySlice, frequencyStruct{word: key, frq: val})
	}
	sort.Slice(frequencySlice, func(i, j int) bool {
		if frequencySlice[i].frq > frequencySlice[j].frq {
			return true
		}
		if frequencySlice[i].frq == frequencySlice[j].frq {
			res := strings.Compare(frequencySlice[i].word, frequencySlice[j].word)
			if res == -1 {
				return true
			}
		}
		return false
	})
	for i, val := range frequencySlice {
		if i == 10 {
			break
		}
		resultSlice = append(resultSlice, val.word)
	}
	return resultSlice
}
