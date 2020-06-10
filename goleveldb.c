#include "goleveldb.h"
#include "_cgo_export.h"

leveldb_comparator_t* go_leveldb_create_comparator() {
		return leveldb_comparator_create(
			NULL,
			go_leveldb_destructor,
			(int (*)(void*, const char* a, size_t alen, const char* b, size_t blen))(go_leveldb_compare),
			(const char* (*)(void*))(go_leveldb_name));

	}