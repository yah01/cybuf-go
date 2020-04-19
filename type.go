package cybuf

import (
	"reflect"
)

type CyBufType int

const (
	CyBufType_Invalid CyBufType = iota
	CyBufType_Nil
	CyBufType_Bool
	CyBufType_Integer
	CyBufType_Float
	CyBufType_String
	CyBufType_Array
	CyBufType_Object
)

func GetInterfaceValueType(v interface{}) CyBufType {
	switch v.(type) {
	case nil:
		return CyBufType_Nil
	case bool:
		return CyBufType_Bool
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return CyBufType_Integer
	case float32, float64:
		return CyBufType_Float
	case []byte, []rune, string:
		return CyBufType_String
	}

	realValue := reflect.ValueOf(v)
	if realValue.Kind() == reflect.Struct {
		return CyBufType_Object
	}

	if realValue.IsZero() || realValue.IsNil() {
		return CyBufType_Nil
	}
	switch realValue.Kind() {
	case reflect.Array, reflect.Slice:
		return CyBufType_Array
	case reflect.Map:
		return CyBufType_Object
	}

	return CyBufType_Invalid
}

func GetReflectValueType(v reflect.Value) CyBufType {
	return GetInterfaceValueType(v.Interface())
}

func GetBytesValueType(v []byte) CyBufType {
	if IsStringValue(v) {
		return CyBufType_String
	}
	if IsArrayValue(v) {
		return CyBufType_Array
	}
	if IsObjectValue(v) {
		return CyBufType_Object
	}
	if IsNilType(v) {
		return CyBufType_Nil
	}
	if IsBoolType(string(v)) {
		return CyBufType_Bool
	}
	if IsIntegerValue(v) {
		return CyBufType_Integer
	}
	if IsFloatValue(v) {
		return CyBufType_Float
	}
	return CyBufType_Invalid
}

func GetBytesValueComplexType(v []byte) CyBufType {
	if IsStringValue(v) {
		return CyBufType_String
	}
	if IsArrayValue(v) {
		return CyBufType_Array
	}
	if IsObjectValue(v) {
		return CyBufType_Object
	}
	return CyBufType_Invalid
}

func GetBytesValueSimpleType(v []byte) CyBufType {
	if IsNilType(v) {
		return CyBufType_Nil
	}
	if IsBoolType(string(v)) {
		return CyBufType_Bool
	}
	if IsIntegerValue(v) {
		return CyBufType_Integer
	}
	if IsFloatValue(v) {
		return CyBufType_Float
	}
	return CyBufType_Invalid
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
