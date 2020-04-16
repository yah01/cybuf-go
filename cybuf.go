package cybuf

import (
	"reflect"
	"strconv"
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
		key       []rune
		value     []rune
		valueType CybufType
	)

	rv := v.(*map[string]interface{})
	for i := 0; i < len(data); {

		//log.Println("nextKey:")
		key, i = nextKey(data, i)
		if key == nil {
			return &ParseError{
				Stage: ParseStage_Key,
				Index: i,
				Char:  data[i],
			}
		}
		// todo check key
		//log.Println("key:", string(key))

		//log.Println("nextColon:")
		i = nextColon(data, i)
		if i == -1 {
			return &ParseError{
				Stage: ParseStage_Colon,
				Index: i,
				Char:  data[i],
			}
		}
		//log.Println("colon:", string(data[i-1]))

		//log.Println("nextValue:")
		value, valueType, i = nextValue(data, i)
		if value == nil {
			return &ParseError{
				Stage: ParseStage_Value,
				Index: i,
				Char:  data[i],
			}
		}

		//log.Println(key, value, valueType)
		switch valueType {
		case CybufType_Nil:
			(*rv)[string(key)] = nil
		case CybufType_Number:
			(*rv)[string(key)], _ = strconv.ParseInt(string(value), 10, 64)
		case CybufType_Decimal:
			(*rv)[string(key)], _ = strconv.ParseFloat(string(value), 64)
		case CybufType_String:
			(*rv)[string(key)] = strings.Trim(string(value), `'"`)
		}
		//log.Println(*rv)
	}

	return nil
}

func nextKey(data []rune, offset int) ([]rune, int) {
	// Find first non-space character
	for offset < len(data) && unicode.IsSpace(data[offset]) {
		offset++
	}
	if offset == len(data) {
		return nil, offset
	}
	start := offset

	// Find key until meet the first space character of colon
	for c := data[offset]; offset < len(data) && !unicode.IsSpace(c) && c != ':'; c = data[offset] {
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
			return i + 1
		}
	}

	return -1
}

func nextValue(data []rune, offset int) ([]rune, CybufType, int) {
	// Find first non-space character
	for offset < len(data) && unicode.IsSpace(data[offset]) {
		offset++
	}
	if offset == len(data) {
		//log.Println("offset == len(data)")
		return nil, CybufType_Nil, offset
	}

	start := offset

	var (
		value     []rune
		valueType CybufType
	)
	switch data[offset] {
	case '{':
		valueType = CybufType_Object
	case '[':
		valueType = CybufType_Array
	case '"', '\'':
		valueType = CybufType_String
	}

	if IsBoundChar(data[offset]) {
		//log.Println("IsBoundChar")
		value, offset = findRightBound(data, offset)
		//log.Println(value, valueType, offset)
		return value, valueType, offset
	} else {
		for offset < len(data) && !unicode.IsSpace(data[offset]) {
			offset++
		}
		value = data[start:offset]
		if _, err := strconv.ParseFloat(string(value), 64); err == nil {
			valueType = CybufType_Decimal
		} else if _, err = strconv.ParseInt(string(value), 10, 64); err == nil {
			valueType = CybufType_Number
		}

		if valueType == CybufType_Nil && string(value) != "nil" {
			return nil, valueType, offset
		}

		return value, valueType, offset
	}
}

// data[offset] must be non-space character
func findRightBound(data []rune, offset int) ([]rune, int) {
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
			//log.Println("data:", data[start:offset+1])
			return data[start : offset+1], offset + 1
		}

		offset++
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
