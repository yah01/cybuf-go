package cybuf

import "unicode"

type CybufType int

const (
	CybufType_Nil CybufType = iota
	CybufType_Number
	CybufType_Decimal
	CybufType_String
	CybufType_Array
	CybufType_Object
)

func IsAllDigit(data []rune) bool {
	for _, c := range data {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func IsBoundChar(c rune) bool {
	switch c {
	case '{', '}', '[', ']', '"', '\'':
		return true
	}
	return false
}

// c must be a bound character
func BoundMap(c rune) rune {
	switch c {
	case '{':
		return '}'
	case '}':
		return '{'
	case '[':
		return ']'
	case ']':
		return '['
	}
	return c
}
