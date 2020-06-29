package gogeardb

// #include "leveldb/c.h"
// #include "gogeardb.h"
import "C"
import (
	"bytes"
)

type Comparator struct {
	c *C.leveldb_comparator_t
}

func NewComparator() *Comparator {
	return &Comparator{c: C.go_leveldb_create_comparator()}
}

func (c *Comparator) Destroy() {
	C.leveldb_comparator_destroy(c.c)
}

//export go_geardb_destructor
func go_geardb_destructor() {}

//export go_geardb_compare
func go_geardb_compare(state int, cKeyA *C.char, cLenA C.size_t, cKeyB *C.char, cLenB C.size_t) int {
	keyA := charToByte(cKeyA, cLenA)
	keyB := charToByte(cKeyB, cLenB)
	return bytes.Compare(keyA, keyB)
}

//export go_geardb_name
func go_geardb_name(state int) *C.char {
	return C.CString("go-leveldb-comparator")
}
