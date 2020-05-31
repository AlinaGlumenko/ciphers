package ciphers

import (
	"regexp"
	"strings"
)

type PigLatin struct {
	Name string
}

var wordEnds []string = []string{"ay", "yay"}

func (pl PigLatin) Encode(text string) (string, bool) {
	sourceText := strings.TrimSpace(text)
	result := sourceText
	r := regexp.MustCompile(`(\w+){1,}`)
	wordsArr := r.FindAllString(sourceText, -1)

	startInd := 0
	for _, word := range wordsArr {
		newWord := strToPigLatin(word)
		result = strings.Replace(result, result[startInd:], strings.Replace(result[startInd:], word, newWord, 1), 1)
		startInd += strings.Index(result[startInd:], newWord) + len(newWord)
	}
	return result, false
}

func isVowel(c byte) bool {
	return (c == byte('A')) || (c == byte('a')) ||
		(c == byte('E')) || (c == byte('e')) ||
		(c == byte('I')) || (c == byte('i')) ||
		(c == byte('O')) || (c == byte('o')) ||
		(c == byte('U')) || (c == byte('u')) ||
		(c == byte('Y')) || (c == byte('y'))
}

func strToPigLatin(word string) string {
	bytes := []byte(word)
	var ind int

	for i, letter := range bytes {
		if isVowel(letter) {
			ind = i
			break
		}
	}

	result := string(bytes[ind:]) + string(bytes[:ind])
	if ind == 0 && isVowel(bytes[ind]) {
		result += wordEnds[1]
	} else {
		result += wordEnds[0]
	}
	return result
}
