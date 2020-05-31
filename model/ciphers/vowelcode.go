package ciphers

import "strings"

type VowelCode struct {
	Name string
}

var vowelsNums map[byte]byte
var reversedMap map[byte]byte

func init() {
	vowelsNums = make(map[byte]byte)
	vowelsNums['a'] = '1'
	vowelsNums['e'] = '2'
	vowelsNums['i'] = '3'
	vowelsNums['o'] = '4'
	vowelsNums['u'] = '5'

	reversedMap = reverseMap(vowelsNums)
}

func (vc VowelCode) Encode(text string) (string, bool) {
	return convert(text, vowelsNums)
}

func (vc VowelCode) Decode(text string) (string, bool) {
	return convert(text, reversedMap)
}

func convert(text string, chars map[byte]byte) (string, bool) {
	text = strings.TrimSpace(text)
	sourceBytes := []byte(text)
	bytes := []byte(strings.ToLower(text))
	for i, letter := range bytes {
		if val, ok := chars[byte(letter)]; ok {
			sourceBytes[i] = val
		}
	}
	return string(sourceBytes), false
}

func reverseMap(m map[byte]byte) map[byte]byte {
	n := make(map[byte]byte)
	for k, v := range m {
		n[v] = k
	}
	return n
}
