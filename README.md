# leveldb-go-wrapper
a simple go wrapper for geardb

## Usage
CGO_CFLAGS="-I/path/to/geardb/include" /\
CGO_LDFLAGS="-L/path/to/geardb/lib -lgeardb" /\
  go get github.com/GlitterIsMe/gogeardb

## Attention
- Some interce has not been implemented, like WriteBatch::Iterator
- Some module has not been tested, like filterpolicy and snappy compression
- Some implementation maybe ineifficient(e.g. memory management with CGO)

