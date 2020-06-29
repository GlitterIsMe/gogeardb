package gogeardb

// #include "stdlib.h"
// #include "leveldb/c.h"
import "C"
import (
	"errors"
	"unsafe"
)

type DB struct {
	db     *C.leveldb_t
	opt    *Options
	dbName string
}

func OpenDB(options *Options, name string) (db *DB, err error) {
	var (
		cErr  *C.char
		cName = C.CString(name)
	)
	defer C.free(unsafe.Pointer(cName))
	cLevelDB := C.leveldb_open(options.c, cName, &cErr)
	if cErr != nil {
		defer C.free(unsafe.Pointer(cErr))
		return nil, errors.New(C.GoString(cErr))
	}
	return &DB{
		db:     cLevelDB,
		opt:    options,
		dbName: name,
	}, nil
}

func (db *DB) Close() {
	C.leveldb_close(db.db)
}

func (db *DB) Put(options *WriteOptions, key, value []byte) error {
	var cErr *C.char
	C.leveldb_put(db.db, options.c, bytesToChar(key), C.size_t(len(key)), bytesToChar(value), C.size_t(len(value)), &cErr)
	if cErr != nil {
		defer C.free(unsafe.Pointer(cErr))
		return errors.New(C.GoString(cErr))
	}
	return nil
}

func (db *DB) Delete(options *WriteOptions, key []byte) error {
	var cErr *C.char
	C.leveldb_delete(db.db, options.c, bytesToChar(key), C.size_t(len(key)), &cErr)
	if cErr != nil {
		defer C.free(unsafe.Pointer(cErr))
		return errors.New(C.GoString(cErr))
	}
	return nil
}

func (db *DB) Get(options *ReadOptions, key []byte) (value []byte, err error) {
	var (
		cValueSize C.size_t
		cValue     *C.char
		cErr       *C.char
	)
	cValue = C.leveldb_get(db.db, options.c, bytesToChar(key), C.size_t(len(key)), &cValueSize, &cErr)
	if cErr != nil {
		defer C.free(unsafe.Pointer(cErr))
		return nil, errors.New(C.GoString(cErr))
	}
	return charToByte(cValue, cValueSize), nil
}

func (db *DB) NewIterator(options *ReadOptions) *Iterator {
	return &Iterator{
		c: C.leveldb_create_iterator(db.db, options.c),
	}
}

func DestroyDB(opt *Options, name string) error {
	var cErr *C.char
	var cDbName *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cDbName))
	C.leveldb_destroy_db(opt.c, cDbName, &cErr)
	if cErr != nil {
		defer C.free(unsafe.Pointer(cErr))
		return errors.New(C.GoString(cErr))
	}
	return nil
}
