package goleveldb_wrapper

// #include "leveldb/c.h"
import "C"

type HmManager struct {
	c *C.geardb_hmmanager_t
}

func NewHmManager() *HmManager {
	return &HmManager{c: C.geardb_hmmanager_create()}
}

func (hm *HmManager) Destroy() {
	C.geardb_hmmanager_destroy(hm.c)
}
