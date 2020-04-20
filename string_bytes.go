package cybuf

import "unsafe"

func string2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	b := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&b))
}

func bytes2string(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}