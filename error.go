package cybuf

import (
	"fmt"
	"reflect"
)

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

type ParseStage string

const (
	ParseStage_Key   ParseStage = "key"
	ParseStage_Colon ParseStage = "colon"
	ParseStage_Value ParseStage = "value"
)

type ParseError struct {
	Stage ParseStage
	Index int
	Char  rune
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("cybuf: Can't parse from %d(%s) when finding %s", e.Index, string(e.Char), e.Stage)
}
