package goleveldb_wrapper

// #include "leveldb/c.h"
import "C"

type Iterator struct {
	c *C.leveldb_iterator_t
}

func (it *Iterator) Valid() bool {
	var v C.uchar = C.leveldb_iter_valid(it.c)
	return charToBoll(v)
}

func (it *Iterator) SeekToFirst() {
	C.leveldb_iter_seek_to_first(it.c)
}

func (it *Iterator) SeekToLast() {
	C.leveldb_iter_seek_to_last(it.c)
}

func (it *Iterator) Seek(k []byte) {
	var cSeekKey *C.char = bytesToChar(k)
	C.leveldb_iter_seek(it.c, cSeekKey, C.size_t(len(k)))
}

func (it *Iterator) Next() {
	C.leveldb_iter_next(it.c)
}

func (it *Iterator) Prev() {
	C.leveldb_iter_prev(it.c)
}

func (it *Iterator) Key() []byte {
	var keyLen C.size_t
	var key *C.char = C.leveldb_iter_key(it.c, &keyLen)
	if key == nil {
		return nil
	}
	return charToByte(key, keyLen)
}

func (it *Iterator) Value() []byte {
	var valueLen C.size_t
	var value *C.char = C.leveldb_iter_value(it.c, &valueLen)
	if value == nil {
		return nil
	}
	return charToByte(value, valueLen)
}

func (it *Iterator) Destroy() {
	C.leveldb_iter_destroy(it.c)
}
