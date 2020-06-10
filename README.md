# leveldb-go-wrapper
a simple go wrapper for leveldb

## Usage
CGO_CFLAGS="-I/path/to/leveldb/include" \
CGO_LDFLAGS="-L/path/to/leveldb/lib -lleveldb" \
  go get github.com/GlitterIsMe/goleveldb_wrapper

## Attention
- Some interce has not been implemented, like WriteBatch::Iterator
- Some module has not been tested, like filterpolicy and snappy compression
- Some implementation maybe ineifficient(e.g. memory management with CGO)

