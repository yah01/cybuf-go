package cybuf

import (
	"reflect"
	"strconv"
)

func Marshal(v interface{}) ([]byte, error) {
	cybufBytes := []byte("{\n")
	bytes, err := marshal(v, 1)
	if err != nil {
		return nil, err
	}

	cybufBytes = append(cybufBytes, bytes...)
	cybufBytes = append(cybufBytes, '}')
	return cybufBytes, nil
}

func marshal(v interface{}, tabCount int) ([]byte, error) {
	var (
		cybufBytes []byte
	)

	rv := v.(map[string]interface{})

	tabsStr := ""
	for i := 0; i < tabCount; i++ {
		tabsStr += "\t"
	}
	tabsBytes := []byte(tabsStr)

	for key, value := range rv {
		var valueBytes []byte
		cybufBytes = append(cybufBytes, tabsBytes...)
		cybufBytes = append(cybufBytes, []byte(key)...)
		cybufBytes = append(cybufBytes, ':', ' ')

		realValue := reflect.ValueOf(value)
		valueType := GetInterfaceValueType(value)
		switch valueType {
		case CybufType_Nil:
			valueBytes = []byte("nil")
		case CybufType_Integer:
			valueBytes = strconv.AppendInt(valueBytes, realValue.Int(), 10)
		case CybufType_Float:
			valueBytes = strconv.AppendFloat(valueBytes, realValue.Float(), 'f', -1, 64)
			//case CybufType_String:
			//	switch realValue.Kind() {
			//	case reflect.String:
			//	case reflect.Slice
			//	}
			//case CybufType_Array:
		}

		cybufBytes = append(cybufBytes, valueBytes...)
		cybufBytes = append(cybufBytes, '\n')
	}

	return cybufBytes, nil
}
