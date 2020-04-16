package cybuf

type CybufType int

const (
	CybufType_nil CybufType = iota
	CybufType_Number
	CybufType_Decimal
	CybufType_String
	CybufType_Array
	CybufType_Object
)

func GetCybufType(value []rune) CybufType {

}
