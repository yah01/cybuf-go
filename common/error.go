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
		return "cybuf: UnmarshalCyBuf(nil)"
	}

	if e.Type.Kind() != reflect.Ptr {
		return "cybuf: UnmarshalCyBuf(non-pointer " + e.Type.String() + ")"
	}
	return "cybuf: UnmarshalCyBuf(nil " + e.Type.String() + ")"
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

type DecodeError string

const (
	DecodeError_Not_Found_Begin DecodeError = "not found beginning('{')"
	DecodeError_Not_Found_End   DecodeError = "not found ending('}')"
)

func (e DecodeError) Error() string {
	return fmt.Sprintf("cybuf: Error happens when decoding: %s", e)
}