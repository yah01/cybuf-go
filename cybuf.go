package cybuf

import (
	"reflect"
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
		key   []rune
		value []rune
	)

	for i := 0; i < len(data); {
		key, i = nextKey(data, i)
		if key == nil {
			return &ParseError{
				Stage: ParseStage_Key,
				Index: i,
				Char:  data[i],
			}
		}

		i = nextColon(data, i)
		if i == -1 {
			return &ParseError{
				Stage: ParseStage_Colon,
				Index: i,
				Char:  data[i],
			}
		}

		value, i = nextValue(data, i)
		if value == nil {
			return &ParseError{
				Stage: ParseStage_Value,
				Index: i,
				Char:  data[i],
			}
		}
	}

	return nil
}

func nextKey(data []rune, offset int) ([]rune, int) {
	var (
		start = offset
	)

	for offset < len(data) && unicode.IsSpace(data[offset]) {
		offset++
	}
	if offset == len(data) {
		return nil, offset
	}

	start = offset

	for c := data[offset]; !unicode.IsSpace(c) && c != ':'; c = data[offset] {
		offset++
	}

	if offset == len(data) {
		return nil, offset
	}

	return data[start:offset], offset
}

func nextColon(data []rune, offset int) int {
	for i := offset; i < len(data); i++ {
		if data[i] == ':' {
			return i
		}
	}

	return -1
}

func nextValue(data []rune, offset int) ([]rune, CybufType, int) {
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
