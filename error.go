package cybuf

import "reflect"

type InvalidUnmarshalError struct {
	Type reflect.Type
}

func (e *InvalidUnmarshalError) Error() string {
	if e.Type == nil {
		return "cybuf: Unmarshal(nil)"
	}

	if e.Type.Kind() != reflect.Ptr {
		return "cybuf: Unmarshal(non-pointer " + e.Type.String() + ")"
	}
	return "cybuf: Unmarshal(nil " + e.Type.String() + ")"
}
