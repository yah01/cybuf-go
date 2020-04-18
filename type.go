package cybuf

import (
	"reflect"
)

type CybufType int

const (
	CybufType_Invalid CybufType = iota
	CybufType_Nil
	CybufType_Bool
	CybufType_Integer
	CybufType_Float
	CybufType_String
	CybufType_Array
	CybufType_Object
)

func GetInterfaceValueType(v interface{}) CybufType {
	switch v.(type) {
	case bool:
		return CybufType_Bool
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return CybufType_Integer
	case float32, float64:
		return CybufType_Float
	case []byte, []rune, string:
		return CybufType_String
	}

	realValue := reflect.ValueOf(v)
	if realValue.Kind() == reflect.Struct {
		return CybufType_Object
	}

	if realValue.IsNil() {
		return CybufType_Nil
	}
	switch realValue.Kind() {
	case reflect.Array, reflect.Slice:
		return CybufType_Array
	case reflect.Map:
		return CybufType_Object
	}

	return CybufType_Invalid
}

func GetBytesValueType(v []byte) CybufType {
	if IsStringValue(v) {
		return CybufType_String
	}
	if IsArrayValue(v) {
		return CybufType_Array
	}
	if IsObjectValue(v) {
		return CybufType_Object
	}
	if IsNilType(v) {
		return CybufType_Nil
	}
	if IsBoolType(string(v)) {
		return CybufType_Bool
	}
	if IsIntegerValue(v) {
		return CybufType_Integer
	}
	if IsFloatValue(v) {
		return CybufType_Float
	}
	return CybufType_Invalid
}

func GetBytesValueComplexType(v []byte) CybufType {
	if IsStringValue(v) {
		return CybufType_String
	}
	if IsArrayValue(v) {
		return CybufType_Array
	}
	if IsObjectValue(v) {
		return CybufType_Object
	}
	return CybufType_Invalid
}

func GetBytesValueSimpleType(v []byte) CybufType {
	if IsNilType(v) {
		return CybufType_Nil
	}
	if IsBoolType(string(v)) {
		return CybufType_Bool
	}
	if IsIntegerValue(v) {
		return CybufType_Integer
	}
	if IsFloatValue(v) {
		return CybufType_Float
	}
	return CybufType_Invalid
}

func IsNilType(v []byte) bool {
	return v[0] == 'n' && v[1] == 'i' && v[2] == 'l'
}

func IsBoolType(v string) bool {
	return v == "true" || v == "True" || v == "false" || v == "False"
}

func IsIntegerValue(v []byte) bool {
	for i := 0; i < len(v); i++ {
		if v[i] < '0' || v[i] > '9' {
			return false
		}
	}
	return true
}

func IsFloatValue(v []byte) bool {
	sawDot := false
	for i := 0; i < len(v); i++ {
		if v[i] == '.' {
			if sawDot {
				return false
			}
			sawDot = true
		} else if v[i] < '0' || v[i] > '9' {
			return false
		}
	}
	return true
}

func IsStringValue(v []byte) bool {
	return v[0] == v[len(v)-1] && (
		v[0] == '\'' || v[0] == '"')
}

func IsArrayValue(v []byte) bool {
	return v[0] == v[len(v)-1] && v[0] == '['
}

func IsObjectValue(v []byte) bool {
	return v[0] == v[len(v)-1] && v[0] == '{'
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
