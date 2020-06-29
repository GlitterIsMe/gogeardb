package gogeardb

// #include "leveldb/c.h"
import "C"

type WriteBatch struct {
	c *C.leveldb_writebatch_t
}

func NewWriteBatch() *WriteBatch {
	return &WriteBatch{c: C.leveldb_writebatch_create()}
}

func (w *WriteBatch) Clear() {
	C.leveldb_writebatch_clear(w.c)
}

func (w *WriteBatch) Put(cKey *C.char, cKeyLen C.size_t, cValue *C.char, cValueLen C.size_t) {
	C.leveldb_writebatch_put(w.c, cKey, cKeyLen, cValue, cValueLen)
}

func (w *WriteBatch) Delete(cKey *C.char, cKeyLen C.size_t) {
	C.leveldb_writebatch_delete(w.c, cKey, cKeyLen)
}

func (w *WriteBatch) Iterator() {

}
