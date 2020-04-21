package cybuf

import (
	"bytes"
	"unicode"
)

func IsValidKeyName(name []byte) bool {
	name = bytes.TrimSpace(name)
	if len(name) == 0 {
		return false
	}

	if len(name) > 0 && !unicode.IsLetter(rune(name[0])) && !(name[0] != '_') {
		return false
	}

	for _, c := range name {
		if !unicode.IsLetter(rune(c)) && unicode.IsDigit(rune(c)) && c != '_' {
			return false
		}
	}
	return true
}

func NextKey(data []byte, offset int) ([]byte, int) {
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

func NextColon(data []byte, offset int) int {
	for i := offset; i < len(data); i++ {
		if data[i] == ':' {
			return i + 1
		}
	}
	return -1
}

func NextValue(data []byte, offset int) ([]byte, CyBufType, int) {
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
		value, offset = FindRightBound(data, offset)
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

func NextKeyValuePair(data []byte, offset int) ([]byte, []byte, CyBufType, int, error) {
	var (
		key       []byte
		value     []byte
		valueType CyBufType
	)

	key, offset = NextKey(data, offset)
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

	offset = NextColon(data, offset)
	if data[offset-1] != ':' {
		return nil, nil, valueType, offset, &ParseError{
			Stage: ParseStage_Colon,
			Index: offset,
			Char:  rune(data[offset]),
		}
	}

	value, valueType, offset = NextValue(data, offset)
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
func FindRightBound(data []byte, offset int) ([]byte, int) {
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
