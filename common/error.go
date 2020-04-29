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
	MarshalStage_Key   MarshalStage = "key"
	MarshalStage_Colon MarshalStage = "colon"
	MarshalStage_Value MarshalStage = "value"

	UnmarshalStage_Key   UnmarshalStage = "key"
	UnmarshalStage_Colon UnmarshalStage = "colon"
	UnmarshalStage_Value UnmarshalStage = "value"

	UnmarshalInfo_NoKey   UnmarshalInfo = "can't parse key"
	UnmarshalInfo_NoColon UnmarshalInfo = "can't parse colon"
	UnmarshalInfo_NoValue UnmarshalInfo = "can't parse value"
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

func NewUnmarshalError(position int, typ CyBufType, stage UnmarshalStage, info UnmarshalInfo) *UnmarshalError {
	return &UnmarshalError{
		Position: position,
		Type:     typ,
		Stage:    stage,
		Info:     info,
	}
}

func (e *UnmarshalError) Error() string {
	return fmt.Sprintf("cybuf: Unmarshal error when parsing %s at %d with type %+v: %s", e.Stage, e.Position, e.Type, e.Info)
}
