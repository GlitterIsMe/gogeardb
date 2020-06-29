package gogeardb

// #include "leveldb/c.h"
import "C"

type Cache struct {
	c *C.leveldb_cache_t
}

func NewCache(size uint) *Cache {
	return &Cache{c: C.leveldb_cache_create_lru(C.size_t(size))}
}

func (c *Cache) Destroy() {
	C.leveldb_cache_destroy(c.c)
}
