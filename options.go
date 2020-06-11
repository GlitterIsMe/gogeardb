package goleveldb_wrapper

// #include "leveldb/c.h"
import "C"

const (
	NoCompression     = int(C.leveldb_no_compression)
	SnappyCompression = int(C.leveldb_snappy_compression)
)

// Options
type Options struct {
	c      *C.leveldb_options_t
	cmp    *Comparator
	cache  *Cache
	env    *Env
	filter *FilterPolicy
}

func NewOptions() *Options {
	return &Options{c: C.leveldb_options_create()}
}

func NewDefaultOptions() *Options {
	opt := NewOptions()
	opt.SetComparator(NewComparator())
	opt.SetCache(NewCache(16 * 1024 * 1024))
	opt.SetEnv(NewDefaultEnv())
	return opt
}

func (opt *Options) SetComparator(cmp *Comparator) {
	opt.cmp = cmp
	C.leveldb_options_set_comparator(opt.c, cmp.c)
}

func (opt *Options) SetFilterPolicy(flt *FilterPolicy) {
	opt.filter = flt
	C.leveldb_options_set_filter_policy(opt.c, flt.c)
}

func (opt *Options) SetCreateIfMissing(create bool) {
	C.leveldb_options_set_create_if_missing(opt.c, boolToChar(create))
}

func (opt *Options) SetErrorIfExists(exists bool) {
	C.leveldb_options_set_error_if_exists(opt.c, boolToChar(exists))
}

func (opt *Options) SetParanoidChecks(paranoid bool) {
	C.leveldb_options_set_paranoid_checks(opt.c, boolToChar(paranoid))
}

func (opt *Options) SetEnv(env *Env) {
	opt.env = env
	C.leveldb_options_set_env(opt.c, env.c)
}

func (opt *Options) SetWriteBufferSize(size uint) {
	C.leveldb_options_set_write_buffer_size(opt.c, C.size_t(size))
}

func (opt *Options) SetMaxOpenFiles(size uint) {
	C.leveldb_options_set_max_open_files(opt.c, C.int(size))
}

func (opt *Options) SetCache(cache *Cache) {
	opt.cache = cache
	C.leveldb_options_set_cache(opt.c, cache.c)
}

func (opt *Options) SetBlockSize(size uint) {
	C.leveldb_options_set_block_size(opt.c, C.size_t(size))
}

func (opt *Options) SetBlockRestartInterval(size uint) {
	C.leveldb_options_set_block_restart_interval(opt.c, C.int(size))
}

func (opt *Options) SetCompression(compression int) {
	C.leveldb_options_set_compression(opt.c, C.int(compression))
}

func (opt *Options) Destroy() {
	C.leveldb_options_destroy(opt.c)
}

func (opt *Options) DestroyDefault() {
	opt.cmp.Destroy()
	opt.cache.Destroy()
	opt.env.Destroy()
	C.leveldb_options_destroy(opt.c)
}

// ReadOptions
type ReadOptions struct {
	c *C.leveldb_readoptions_t
}

func NewReadOptions() *ReadOptions {
	return &ReadOptions{c: C.leveldb_readoptions_create()}
}

func (opt *ReadOptions) SetVerifyChecksums(check bool) {
	C.leveldb_readoptions_set_verify_checksums(opt.c, boolToChar(check))
}

func (opt *ReadOptions) SetFillCache(fill bool) {
	C.leveldb_readoptions_set_fill_cache(opt.c, boolToChar(fill))
}

func (opt *ReadOptions) SetSnapshot(snap Snapshot) {
	C.leveldb_readoptions_set_snapshot(opt.c, snap.c)
}

func (opt *ReadOptions) Destroy() {
	C.leveldb_readoptions_destroy(opt.c)
}

// WriteOptions
type WriteOptions struct {
	c *C.leveldb_writeoptions_t
}

func NewWriteOptions() *WriteOptions {
	return &WriteOptions{c: C.leveldb_writeoptions_create()}
}

func (opt *WriteOptions) SetSync(sync bool) {
	C.leveldb_writeoptions_set_sync(opt.c, boolToChar(sync))
}

func (opt *WriteOptions) Destroy() {
	C.leveldb_writeoptions_destroy(opt.c)
}
