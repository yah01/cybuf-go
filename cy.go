package cybuf

import (
	"reflect"
	"strings"
	"unicode"
)

func Unmarshal(data []byte, v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		return &InvalidUnmarshalError{
			Type: reflect.TypeOf(v),
		}
	}

	realData := []rune(string(data))

	return unmarshal(realData, v)
}

func unmarshal(data []rune, v interface{}) error {
	var (
		FindKey   bool
		FindValue bool
		InString  bool
		InMap     bool
		InArray   bool
	)

	for i := 0; i < len(data); i++ {
		c := data[i]

		switch c {
		case ''
		}
	}

	return nil
}

func nextKey(data []rune, offset int) ([]rune, int) {
	var (
		start = offset

		sawFirstChar bool
	)
	for i, c := range data {

		if sawFirstChar {
			if unicode.IsSpace(c) || c == ':' {
				return data[start:i], i
			}
		} else {
			if unicode.IsSpace(c) {
				continue
			} else {
				start = i
				sawFirstChar = true
			}
		}
	}

	return nil, offset
}

func nextColon(data []rune, offset int) int {
	for i := offset; i < len(data); i++ {
		if data[i] == ':' {
			return i
		}
	}

	return -1
}

func nextValue(data []rune,offset int) ([]rune,int) {

}

func IsValieKeyName(name []rune) bool {
	if len(name) > 0 && !unicode.IsLetter(name[0]) && !(name[0] != '_') {
		return false
	}
	for _, c := range name {
		if !unicode.IsLetter(c) && unicode.IsDigit(c) && c != '_' {
			return false
		}
	}
	return true
}
