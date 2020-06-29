package gogeardb

// #include "leveldb/c.h"
import "C"

type FilterPolicy struct {
	c *C.leveldb_filterpolicy_t
}

func NewBloomFilterPolicy(bitPerKey int) *FilterPolicy {
	return &FilterPolicy{c: C.leveldb_filterpolicy_create_bloom(C.int(bitPerKey))}
}

func (f *FilterPolicy) Destroy() {
	C.leveldb_filterpolicy_destroy(f.c)
}
