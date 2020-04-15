package cybuf

type CybufType int
const (
	CybufType_Number CybufType = iota
	CybufType_Decimal
	CybufType_String
	CybufType_Array
	CybufType_Object
)

func GetCybufType(value []rune) CybufType {

}