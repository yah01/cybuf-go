package cybuf

import (
	"fmt"
	"reflect"
)

type (
	MarshalStage   string
	UnmarshalStage string
	MarshalInfo    string
	UnmarshalInfo  string
)

const (
	ParseStage_Key   MarshalStage = "key"
	ParseStage_Colon MarshalStage = "colon"
	ParseStage_Value MarshalStage = "value"
)

type MarshalError struct {
	Position int
	Type     reflect.Type
	Stage    MarshalStage
	Info     MarshalInfo
}

func NewMarshalError(position int, typ reflect.Type, stage MarshalStage, info MarshalInfo) *MarshalError {
	return &MarshalError{
		Position: position,
		Type:     typ,
		Stage:    stage,
		Info:     info,
	}
}

func (e *MarshalError) Error() string {
	return fmt.Sprintf("cybuf: Marshal error when parsing %s at %d with type %+v: %s", e.Stage, e.Position, e.Type, e.Info)
}

type UnmarshalError struct {
	Position int
	Type     CyBufType
	Stage    UnmarshalStage
	Info     UnmarshalInfo
}
