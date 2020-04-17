package cybuf

import (
	"reflect"
	"unicode"
)

type CybufType int

const (
	CybufType_Nil CybufType = iota
	CybufType_Integer
	CybufType_Float
	CybufType_Bool
	CybufType_String
	CybufType_Array
	CybufType_Object
)

func GetValueType(v interface{}) CybufType {
	switch v.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return CybufType_Integer
	case float32, float64:
		return CybufType_Float
	case []byte, []rune, string:
		return CybufType_String
	}

	realValue := reflect.ValueOf(v)
	if realValue.IsNil() {
		return CybufType_Nil
	}
	switch realValue.Kind() {
	case reflect.Array, reflect.Slice:
		return CybufType_Array
	case reflect.Map, reflect.Struct:
		return CybufType_Object
	}

	return CybufType_Nil
}

func IsAllDigit(data []rune) bool {
	for _, c := range data {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func IsBoundChar(c byte) bool {
	switch c {
	case '{', '}', '[', ']', '"', '\'':
		return true
	}
	return false
}

// c must be a bound character
func BoundMap(c byte) byte {
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
