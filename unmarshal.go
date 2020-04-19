package cybuf

import (
	"bytes"
	"reflect"
	"strconv"
	"unicode"
)

type Unmarshaler interface {
	Unmarshal(data []byte, v interface{}) error
}

func Unmarshal(data []byte, v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		return &InvalidUnmarshalError{
			Type: reflect.TypeOf(v),
		}
	}

	if rv.Elem().Kind() == reflect.Map {
		return unmarshal(data, v)
	} else {
		objMap := new(map[string]interface{})
		err := unmarshal(data, objMap)
		if err != nil {
			return err
		}
		return mapToStruct(objMap, v)
	}
}

func unmarshal(data []byte, v interface{}) error {
	var (
		key       []byte
		keyStr    string
		value     []byte
		valueStr  string
		valueType CyBufType
		err       error
	)

	data = bytes.TrimSpace(data)
	for data[0] == '{' && data[len(data)-1] == '}' {
		data = data[1 : len(data)-1]
	}

	// debugLog.Println("unmarshal data:", string(data))

	rv := v.(*map[string]interface{})
	for i := 0; i < len(data); {

		key, value, valueType, i, err = nextKeyValuePair(data, i)
		if err != nil {
			// errorLog.Println(err)
			return err
		}

		if key == nil && i == len(data) {
			break
		}
		keyStr = string(key)
		//if !IsValidKeyName(key) {
		//	return &ParseError{
		//		Stage: ParseStage_Key,
		//		Index: i,
		//		//Char:  data[i],
		//	}
		//}
		valueStr = string(value)

		// debugLog.Println("value: "+string(value)+", valueType:", valueType)
		switch valueType {
		case CyBufType_Nil:
			(*rv)[keyStr] = nil
		case CyBufType_Integer:
			(*rv)[keyStr], _ = strconv.ParseInt(valueStr, 10, 64)
		case CyBufType_Float:
			(*rv)[keyStr], _ = strconv.ParseFloat(valueStr, 64)
		case CyBufType_String:
			(*rv)[keyStr] = string(value[1 : len(value)-1])
		case CyBufType_Array:
			array := make([]interface{}, 0)
			//waitSub.Add(1)
			err := unmarshalArray(value, &array)
			if err != nil {
				// errorLog.Println(err)
				return err
			}
			(*rv)[keyStr] = array

		case CyBufType_Object:
			var object = make(map[string]interface{})
			err := unmarshal(value, &object)
			if err != nil {
				// errorLog.Println(err)
				return err
			}
			// debugLog.Println(object)
			(*rv)[keyStr] = object
		}

		//debugLog.Println("parsed:", keyStr, valueStr)
	}

	return nil
}

func unmarshalArray(data []byte, v *[]interface{}) error {
	var (
		value     []byte
		valueStr  string
		valueType CyBufType
		realValue interface{}
	)

	data = bytes.TrimSpace(data)
	//debugLog.Println(string(data))
	data = data[1 : len(data)-1]

	//debugLog.Println(string(data))

	// debugLog.Println("unmarshal array data:", string(data))

	for i := 0; i < len(data); {
		value, valueType, i = nextValue(data, i)

		if value == nil {
			if i >= len(data) {
				break
			}
			return &ParseError{
				Stage: ParseStage_Value,
				Index: i,
				//Char:  rune(data[i]),
			}
		}
		valueStr = string(value)

		switch valueType {
		case CyBufType_Nil:
			realValue = nil
		case CyBufType_Bool:
			switch valueStr {
			case "true", "True":
				realValue = true
			case "false", "False":
				realValue = false
			}
		case CyBufType_Integer:
			realValue, _ = strconv.ParseInt(valueStr, 10, 64)
			//realValue = intValue
		case CyBufType_Float:
			realValue, _ = strconv.ParseFloat(valueStr, 64)
		case CyBufType_String:
			realValue = string(value[1 : len(value)-1])
		case CyBufType_Array:
			array := make([]interface{}, 0)
			err := unmarshalArray(value, &array)
			if err != nil {
				// errorLog.Println(err)
				return err
			}
			realValue = array
		case CyBufType_Object:
			var object = make(map[string]interface{})
			err := unmarshal(value, &object)
			if err != nil {
				// errorLog.Println(err)
				return err
			}
			// debugLog.Println(object)
			realValue = object
		}

		// debugLog.Println("append:", realValue)
		*v = append(*v, realValue)
	}

	return nil
}

func nextKey(data []byte, offset int) ([]byte, int) {
	// Find first non-space character
	for offset < len(data) && unicode.IsSpace(rune(data[offset])) {
		offset++
	}
	if offset == len(data) {
		return nil, offset
	}
	start := offset

	// Find key until meet the first space character of colon
	for c := data[offset]; offset < len(data) && !unicode.IsSpace(rune(c)) && c != ':'; c = data[offset] {
		offset++
	}
	if offset == len(data) {
		return nil, offset
	}

	return data[start:offset], offset
}

func nextColon(data []byte, offset int) int {
	for i := offset; i < len(data); i++ {
		if data[i] == ':' {
			return i + 1
		}
	}
	return -1
}

func nextValue(data []byte, offset int) ([]byte, CyBufType, int) {
	// Find first non-space character
	for offset < len(data) && unicode.IsSpace(rune(data[offset])) {
		offset++
	}
	if offset == len(data) {
		return nil, CyBufType_Nil, offset
	}

	start := offset

	var (
		value     []byte
		valueType CyBufType
	)
	switch data[offset] {
	case '{':
		valueType = CyBufType_Object
	case '[':
		valueType = CyBufType_Array
	case '"', '\'':
		valueType = CyBufType_String
	}

	if IsBoundChar(data[offset]) {
		// debugLog.Println("IsBoundChar, offset =", offset, "data = ", string(data[offset:]))
		value, offset = findRightBound(data, offset)
		// debugLog.Println("new offset =", offset)
		return value, valueType, offset
	} else {

		for offset < len(data) && !unicode.IsSpace(rune(data[offset])) {
			offset++
		}
		value = data[start:offset]

		valueType = GetBytesValueSimpleType(value)
		if valueType == CyBufType_Invalid {
			return nil, valueType, offset
		}

		return value, valueType, offset
	}
}

func nextKeyValuePair(data []byte, offset int) ([]byte, []byte, CyBufType, int, error) {
	var (
		key       []byte
		value     []byte
		valueType CyBufType
	)

	key, offset = nextKey(data, offset)
	if key == nil {
		if offset == len(data) {
			return nil, nil, valueType, offset, nil
		}
		return nil, nil, valueType, offset, &ParseError{
			Stage: ParseStage_Key,
			Index: offset,
			Char:  rune(data[offset]),
		}
	}
	// debugLog.Println("key:", string(key))

	offset = nextColon(data, offset)
	if data[offset-1] != ':' {
		return nil, nil, valueType, offset, &ParseError{
			Stage: ParseStage_Colon,
			Index: offset,
			Char:  rune(data[offset]),
		}
	}

	value, valueType, offset = nextValue(data, offset)
	if value == nil {
		return nil, nil, valueType, offset, &ParseError{
			Stage: ParseStage_Value,
			Index: offset,
			//Char:  data[offset],
		}
	}
	// debugLog.Println("value:", string(value))

	return key, value, valueType, offset, nil
}

// data[offset] must be non-space character
func findRightBound(data []byte, offset int) ([]byte, int) {
	var (
		leftBound      = data[offset]
		rightBound     = BoundMap(leftBound)
		leftBoundCount = 1
		start          = offset
	)

	offset++
	for offset < len(data) {
		switch data[offset] {
		// Check case data[offset]==rightBound first, quotes won't work if reverse
		case rightBound:
			leftBoundCount--
		case leftBound:
			leftBoundCount++
		}

		if leftBoundCount == 0 {
			return data[start : offset+1], offset + 1
		}

		offset++
	}

	// debugLog.Println("right bound not found!")
	return nil, offset
}
