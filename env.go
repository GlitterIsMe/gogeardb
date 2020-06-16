package goleveldb_wrapper

// #include "leveldb/c.h"
import "C"

type Env struct {
	c *C.leveldb_env_t
}

func NewDefaultEnv() *Env {
	return &Env{c: C.leveldb_create_default_env()}
}

func (e *Env) Destroy() {
	C.leveldb_env_destroy(e.c)
}

func NewDefaultHmEnv(hm *HmManager) *Env {
	return &Env{
		c: C.geardb_hm_env_create(hm.c),
	}
}
