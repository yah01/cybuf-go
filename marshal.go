package cybuf

import (
	"bytes"
	"reflect"
	"strconv"
)

func Marshal(v interface{}) ([]byte, error) {
	return marshal(v)
}

func marshal(v interface{}) ([]byte, error) {
	var (
		cybufBytes = []byte{'{'}
		realValue  reflect.Value
		valueType  CyBufType

		valueBytes = make([]byte, 0)
	)

	rv := v.(map[string]interface{})

	for key, value := range rv {
		cybufBytes = append(cybufBytes, []byte(key)...)
		cybufBytes = append(cybufBytes, ':')

		valueBytes = valueBytes[:0]
		realValue = reflect.ValueOf(value)
		valueType = GetInterfaceValueType(value)
		switch valueType {
		case CyBufType_Nil:
			valueBytes = []byte("nil")
		case CyBufType_Bool:
			valueBytes = strconv.AppendBool(valueBytes, realValue.Bool())
		case CyBufType_Integer:
			valueBytes = strconv.AppendInt(valueBytes, realValue.Int(), 10)
		case CyBufType_Float:
			valueBytes = strconv.AppendFloat(valueBytes, realValue.Float(), 'f', -1, 64)
		case CyBufType_String:
			switch realValue.Kind() {
			case reflect.String:
				valueBytes = strconv.AppendQuote(valueBytes, realValue.String())
			case reflect.Slice:
				valueBytes = append(valueBytes, realValue.Bytes()...)
			}
		case CyBufType_Array:
			arrayBytes, err := marshalArray(value)
			if err != nil {
				return nil, err
			}
			valueBytes = append(valueBytes, arrayBytes...)
		case CyBufType_Object:
			objectBytes, err := marshal(value)
			if err != nil {
				return nil, err
			}
			valueBytes = append(valueBytes, objectBytes...)
		}
		cybufBytes = append(cybufBytes, valueBytes...)
		if valueType != CyBufType_String && valueType != CyBufType_Array && valueType != CyBufType_Object {
			cybufBytes = append(cybufBytes, MarshalSep)
		}
	}
	if cybufBytes[len(cybufBytes)-1] == MarshalSep {
		cybufBytes[len(cybufBytes)-1] = '}'
	} else {
		cybufBytes = append(cybufBytes, '}')
	}

	return cybufBytes, nil
}

func marshalArray(v interface{}) ([]byte, error) {
	var (
		cybufBytes = []byte{'['}
		value      interface{}
		realValue  reflect.Value
		valueType  CyBufType
		valueBytes = make([]byte, 0)
	)

	rv := reflect.ValueOf(v)

	for i := 0; i < rv.Len(); i++ {
		value = rv.Index(i).Interface()
		realValue = reflect.ValueOf(value)
		valueType = GetInterfaceValueType(value)
		valueBytes = valueBytes[:0]

		switch valueType {
		case CyBufType_Nil:
			valueBytes = []byte("nil")
		case CyBufType_Bool:
			valueBytes = strconv.AppendBool(valueBytes, realValue.Bool())
		case CyBufType_Integer:
			valueBytes = strconv.AppendInt(valueBytes, realValue.Int(), 10)
		case CyBufType_Float:
			valueBytes = strconv.AppendFloat(valueBytes, realValue.Float(), 'f', -1, 64)
		case CyBufType_String:
			switch realValue.Kind() {
			case reflect.String:
				valueBytes = strconv.AppendQuote(valueBytes, realValue.String())
			case reflect.Slice:
				valueBytes = append(valueBytes, realValue.Bytes()...)
			}
		case CyBufType_Array:
			arrayBytes, err := marshalArray(value)
			if err != nil {
				return nil, err
			}
			valueBytes = append(valueBytes, arrayBytes...)
		case CyBufType_Object:
			objectBytes, err := marshal(value)
			if err != nil {
				return nil, err
			}
			valueBytes = append(valueBytes, objectBytes...)
		}
		cybufBytes = append(cybufBytes, valueBytes...)
		if valueType != CyBufType_String && valueType != CyBufType_Array && valueType != CyBufType_Object {
			cybufBytes = append(cybufBytes, MarshalSep)
		}
	}
	if cybufBytes[len(cybufBytes)-1] == MarshalSep {
		cybufBytes[len(cybufBytes)-1] = ']'
	} else {
		cybufBytes = append(cybufBytes, ']')
	}

	return cybufBytes, nil
}

func MarshalIndent(v interface{}) ([]byte, error) {
	return marshalIndent(v, 1)
}

func marshalIndent(v interface{}, tabCount int) ([]byte, error) {
	var (
		cybufBytes = []byte{'{', '\n'}
		realValue  reflect.Value
		valueType  CyBufType

		valueBytes = make([]byte, 0)
	)

	rv := v.(map[string]interface{})

	tabs := bytes.Repeat([]byte{'\t'}, tabCount)
	if tabCount == 0 {
		tabs = []byte{'\t'}
	}

	for key, value := range rv {
		cybufBytes = append(cybufBytes, tabs...)
		cybufBytes = append(cybufBytes, []byte(key)...)
		cybufBytes = append(cybufBytes, ':', ' ')

		valueBytes = valueBytes[:0]
		realValue = reflect.ValueOf(value)
		valueType = GetInterfaceValueType(value)
		switch valueType {
		case CyBufType_Nil:
			valueBytes = []byte("nil")
		case CyBufType_Bool:
			valueBytes = strconv.AppendBool(valueBytes, realValue.Bool())
		case CyBufType_Integer:
			valueBytes = strconv.AppendInt(valueBytes, realValue.Int(), 10)
		case CyBufType_Float:
			valueBytes = strconv.AppendFloat(valueBytes, realValue.Float(), 'f', -1, 64)
		case CyBufType_String:
			switch realValue.Kind() {
			case reflect.String:
				valueBytes = strconv.AppendQuote(valueBytes, realValue.String())
			case reflect.Slice:
				valueBytes = append(valueBytes, realValue.Bytes()...)
			}
		case CyBufType_Array:
			arrayBytes, err := marshalArrayIndent(value, tabCount+1)
			if err != nil {
				return nil, err
			}
			valueBytes = append(valueBytes, arrayBytes...)
		case CyBufType_Object:
			objectBytes, err := marshalIndent(value, tabCount+1)
			if err != nil {
				return nil, err
			}
			valueBytes = append(valueBytes, objectBytes...)
		}
		cybufBytes = append(cybufBytes, valueBytes...)
		cybufBytes = append(cybufBytes, '\n')
	}

	cybufBytes = append(cybufBytes, tabs[1:]...)
	cybufBytes = append(cybufBytes, '}')

	return cybufBytes, nil
}

func marshalArrayIndent(v interface{}, tabCount int) ([]byte, error) {
	var (
		cybufBytes = []byte{'[', '\n'}
		value      interface{}
		realValue  reflect.Value
		valueType  CyBufType
		valueBytes = make([]byte, 0)
	)

	rv := reflect.ValueOf(v)

	tabs := bytes.Repeat([]byte{'\t'}, tabCount)
	if tabCount == 0 {
		tabs = []byte{'\t'}
	}

	for i := 0; i < rv.Len(); i++ {
		cybufBytes = append(cybufBytes, tabs...)

		value = rv.Index(i).Interface()
		realValue = reflect.ValueOf(value)
		valueType = GetInterfaceValueType(value)
		valueBytes = valueBytes[:0]

		switch valueType {
		case CyBufType_Nil:
			valueBytes = []byte("nil")
		case CyBufType_Bool:
			valueBytes = strconv.AppendBool(valueBytes, realValue.Bool())
		case CyBufType_Integer:
			valueBytes = strconv.AppendInt(valueBytes, realValue.Int(), 10)
		case CyBufType_Float:
			valueBytes = strconv.AppendFloat(valueBytes, realValue.Float(), 'f', -1, 64)
		case CyBufType_String:
			switch realValue.Kind() {
			case reflect.String:
				valueBytes = strconv.AppendQuote(valueBytes, realValue.String())
			case reflect.Slice:
				valueBytes = append(valueBytes, realValue.Bytes()...)
			}
		case CyBufType_Array:
			arrayBytes, err := marshalArrayIndent(value, tabCount+1)
			if err != nil {
				return nil, err
			}
			valueBytes = append(valueBytes, arrayBytes...)
		case CyBufType_Object:
			objectBytes, err := marshalIndent(value, tabCount+1)
			if err != nil {
				return nil, err
			}
			valueBytes = append(valueBytes, objectBytes...)
		}
		cybufBytes = append(cybufBytes, valueBytes...)
		cybufBytes = append(cybufBytes, '\n')
	}

	cybufBytes = append(cybufBytes, tabs[1:]...)
	cybufBytes = append(cybufBytes, ']')

	return cybufBytes, nil
}
