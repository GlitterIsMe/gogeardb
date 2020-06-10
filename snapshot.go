package goleveldb_wrapper

// #include "leveldb/c.h"
import "C"

type Snapshot struct {
	c *C.leveldb_snapshot_t
}

func NewSnapshot(db *DB) *Snapshot {
	return &Snapshot{c: C.leveldb_create_snapshot(db.db)}
}

func (s *Snapshot) ReleaseSnapshot(db *DB) {
	C.leveldb_release_snapshot(db.db, s.c)
}
