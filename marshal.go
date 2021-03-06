package cybuf

import (
	"bytes"
	. "github.com/yah01/cybuf-go/common"
	"io/ioutil"
	"reflect"
	"strconv"
)

type Marshaler interface {
	MarshalCyBuf() ([]byte, error)
}

func Marshal(v interface{}) ([]byte, error) {
	kind := reflect.TypeOf(v).Kind()
	if kind == reflect.Map {
		return marshal(v)
	} else if kind == reflect.Struct {
		return marshalStruct(v)
	} else if kind == reflect.Slice || kind == reflect.Array {
		return marshalArray(v)
	}

	return nil, nil
}

func Save(v interface{}, fileName string) error {
	cybufBytes, err := Marshal(v)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, cybufBytes, 0644)
}

func marshal(v interface{}) ([]byte, error) {
	var (
		cybufBytes = []byte{'{'}
		realValue  reflect.Value
		valueType  CyBufType

		valueBytes = make([]byte, 0)
	)
	//log.Println("marshal()")
	rv := v.(map[string]interface{})

	for key, value := range rv {
		cybufBytes = append(cybufBytes, String2bytes(key)...)
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
			if IsSignedInteger(realValue) {
				valueBytes = strconv.AppendInt(valueBytes, realValue.Int(), 10)
			} else {
				valueBytes = strconv.AppendUint(valueBytes, realValue.Uint(), 10)
			}
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
			var (
				objectBytes []byte
				err         error
			)
			if realValue.Kind() == reflect.Map {
				objectBytes, err = marshal(realValue.Interface())
			} else {
				objectBytes, err = marshalStruct(realValue.Interface())
			}
			if err != nil {
				return nil, err
			}
			valueBytes = append(valueBytes, objectBytes...)
		}

		cybufBytes = append(cybufBytes, valueBytes...)
		if valueType != CyBufType_String && valueType != CyBufType_Array && valueType != CyBufType_Object {
			cybufBytes = append(cybufBytes, marshalSep)
		}
	}
	if cybufBytes[len(cybufBytes)-1] == marshalSep {
		cybufBytes[len(cybufBytes)-1] = '}'
	} else {
		cybufBytes = append(cybufBytes, '}')
	}

	return cybufBytes, nil
}

func marshalStruct(v interface{}) ([]byte, error) {
	var (
		cybufBytes = []byte{'{'}
		valueBytes = make([]byte, 0)

		realValue reflect.Value
		valueType CyBufType
		err       error
	)

	//log.Println("marshalStruct()")

	typeValue := reflect.TypeOf(v)
	rv := reflect.ValueOf(v)

	for i := 0; i < rv.NumField(); i++ {
		cybufBytes = append(cybufBytes, String2bytes(typeValue.Field(i).Name)...)
		cybufBytes = append(cybufBytes, ':')

		realValue = rv.Field(i)
		valueType = GetInterfaceValueType(realValue.Interface())

		valueBytes = valueBytes[:0]
		valueBytes, err = handleMarshal(valueType, valueBytes, realValue)
		if err != nil {
			return nil, err
		}

		cybufBytes = append(cybufBytes, valueBytes...)
		if valueType != CyBufType_String && valueType != CyBufType_Array && valueType != CyBufType_Object {
			cybufBytes = append(cybufBytes, marshalSep)
		}
	}
	if cybufBytes[len(cybufBytes)-1] == marshalSep {
		cybufBytes[len(cybufBytes)-1] = '}'
	} else {
		cybufBytes = append(cybufBytes, '}')
	}

	return cybufBytes, nil
}

func marshalArray(v interface{}) ([]byte, error) {
	var (
		cybufBytes = []byte{'['}
		valueBytes = make([]byte, 0)

		realValue reflect.Value
		valueType CyBufType
		err       error
	)

	//log.Println("marshalArray")

	rv := reflect.ValueOf(v)

	for i := 0; i < rv.Len(); i++ {
		realValue = rv.Index(i)
		valueType = GetInterfaceValueType(realValue.Interface())

		valueBytes = valueBytes[:0]
		valueBytes, err = handleMarshal(valueType, valueBytes, realValue)
		if err != nil {
			return nil, err
		}

		cybufBytes = append(cybufBytes, valueBytes...)
		if valueType != CyBufType_String && valueType != CyBufType_Array && valueType != CyBufType_Object {
			cybufBytes = append(cybufBytes, marshalSep)
		}
	}
	if cybufBytes[len(cybufBytes)-1] == marshalSep {
		cybufBytes[len(cybufBytes)-1] = ']'
	} else {
		cybufBytes = append(cybufBytes, ']')
	}

	return cybufBytes, nil
}

func handleMarshal(valueType CyBufType, valueBytes []byte, realValue reflect.Value) ([]byte, error) {

	switch valueType {
	case CyBufType_Nil:
		valueBytes = []byte("nil")
	case CyBufType_Bool:
		valueBytes = strconv.AppendBool(valueBytes, realValue.Bool())
	case CyBufType_Integer:
		if IsSignedInteger(realValue) {
			valueBytes = strconv.AppendInt(valueBytes, realValue.Int(), 10)
		} else {
			valueBytes = strconv.AppendUint(valueBytes, realValue.Uint(), 10)
		}
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
		arrayBytes, err := marshalArray(realValue.Interface())
		if err != nil {
			return nil, err
		}
		valueBytes = append(valueBytes, arrayBytes...)
	case CyBufType_Object:
		var (
			objectBytes []byte
			err         error
		)
		//log.Println("realValue.Kind():",realValue.Type().Kind().String())
		if realValue.Kind() == reflect.Map || realValue.Kind() == reflect.Interface {
			objectBytes, err = marshal(realValue.Interface())
		} else {
			objectBytes, err = marshalStruct(realValue.Interface())
		}
		if err != nil {
			return nil, err
		}
		valueBytes = append(valueBytes, objectBytes...)
	}

	return valueBytes, nil
}

func MarshalIndent(v interface{}) ([]byte, error) {
	kind := reflect.TypeOf(v).Kind()
	if kind == reflect.Map {
		return marshalIndent(v, 1)
	} else if kind == reflect.Struct {
		return marshalStructIndent(v, 1)
	}

	return nil, nil
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

	for key, value := range rv {
		cybufBytes = append(cybufBytes, tabs...)
		cybufBytes = append(cybufBytes, String2bytes(key)...)
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

func marshalStructIndent(v interface{}, tabCount int) ([]byte, error) {
	var (
		cybufBytes = []byte{'{', '\n'}
		realValue  reflect.Value
		valueType  CyBufType

		valueBytes = make([]byte, 0)
	)

	typeValue := reflect.TypeOf(v)
	rv := reflect.ValueOf(v)

	tabs := bytes.Repeat([]byte{'\t'}, tabCount)

	for i := 0; i < rv.NumField(); i++ {
		field := typeValue.Field(i)

		cybufBytes = append(cybufBytes, tabs...)
		cybufBytes = append(cybufBytes, String2bytes(field.Name)...)
		cybufBytes = append(cybufBytes, ':', ' ')

		valueBytes = valueBytes[:0]
		realValue = rv.Field(i)
		valueType = GetInterfaceValueType(realValue.Interface())
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
			arrayBytes, err := marshalArrayIndent(realValue.Interface(), tabCount+1)
			if err != nil {
				return nil, err
			}
			valueBytes = append(valueBytes, arrayBytes...)
		case CyBufType_Object:
			var (
				objectBytes []byte
				err         error
			)
			if realValue.Kind() == reflect.Map {
				objectBytes, err = marshalIndent(realValue.Interface(), tabCount+1)
			} else {
				objectBytes, err = marshalStructIndent(realValue.Interface(), tabCount+1)
			}
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
			var (
				objectBytes []byte
				err         error
			)
			if realValue.Kind() == reflect.Map {
				objectBytes, err = marshalIndent(realValue.Interface(), tabCount+1)
			} else {
				objectBytes, err = marshalStructIndent(realValue.Interface(), tabCount+1)
			}
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
