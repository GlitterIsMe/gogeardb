#include "gogeardb.h"
#include "_cgo_export.h"

leveldb_comparator_t* go_geardb_create_comparator() {
		return leveldb_comparator_create(
			NULL,
			go_geardb_destructor,
			(int (*)(void*, const char* a, size_t alen, const char* b, size_t blen))(go_geardb_compare),
			(const char* (*)(void*))(go_geardb_name));

	}