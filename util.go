package gogeardb

import (
	"C"
	"reflect"
	"unsafe"
)

func charToByte(data *C.char, len C.size_t) []byte {
	var value []byte
	var sh = (*reflect.SliceHeader)(unsafe.Pointer(&value))
	sh.Cap, sh.Len, sh.Data = int(len), int(len), uintptr(unsafe.Pointer(data))
	return value
}

func boolToChar(b bool) C.uchar {
	if b {
		return 1
	}
	return 0
}

func charToBoll(c C.uchar) bool {
	if c == 0 {
		return false
	}
	return true
}

func bytesToChar(k []byte) *C.char {
	var c *C.char
	if len(k) > 0 {
		c = (*C.char)(unsafe.Pointer(&k[0]))
	}
	return c
}
