package gogeardb

import (
	"fmt"
	"strconv"
	"testing"
)

const dbName = "go-leveldb"

func TestDB(t *testing.T) {
	opt := NewOptions()
	defer opt.Destroy()

	comparator := NewComparator()
	defer comparator.Destroy()

	cache := NewCache(16 * 1024 * 1024)
	defer cache.Destroy()

	hm := NewHmManager()
	defer hm.Destroy()

	env := NewDefaultHmEnv(hm)
	defer env.Destroy()

	opt.SetComparator(comparator)
	opt.SetCreateIfMissing(true)
	opt.SetErrorIfExists(true)
	opt.SetParanoidChecks(false)
	opt.SetEnv(env)
	opt.SetWriteBufferSize(4 << 20)
	opt.SetMaxOpenFiles(1000)
	opt.SetCache(cache)
	opt.SetBlockSize(4096)
	opt.SetBlockRestartInterval(16)
	opt.SetCompression(SnappyCompression)
	opt.SetHmManager(hm)

	destroyDB := func() {
		err := DestroyDB(opt, dbName)
		if err != nil {
			t.Error("close failed", err.Error())
		} else {
			fmt.Println("close go-leveldb")
		}
	}
	destroyDB()

	// Test Open
	db, err := OpenDB(opt, dbName)

	if err != nil {
		t.Error("open failed", err.Error())
		return
	} else {
		fmt.Println("open db go-leveldb success")
	}

	writeOptions := NewWriteOptions()
	writeOptions.SetSync(false)

	// Test Put
	put_err := db.Put(writeOptions, []byte("foo0"), []byte("bar0"))
	if put_err != nil {
		t.Error("put failed", put_err.Error())
	} else {
		fmt.Println("put success")
	}

	rOption := NewReadOptions()
	rOption.SetFillCache(true)
	rOption.SetVerifyChecksums(true)

	// Test Get
	value, get_err := db.Get(rOption, []byte("foo0"))
	if get_err != nil {
		t.Error("get failed", get_err.Error())
	} else if len(value) == 0 {
		t.Error("not found")
	} else {
		fmt.Println("get", string(value))
	}

	var iterErr error
	for i := 0; i < 10; i++ {
		iterErr = db.Put(writeOptions, []byte("foo"+strconv.Itoa(i)), []byte("bar"+strconv.Itoa(i)))
		fmt.Println("put", "foo"+strconv.Itoa(i), "bar"+strconv.Itoa(i))
		if iterErr != nil {
			t.Error("put failed", err.Error())
		}
	}

	_ = db.Put(writeOptions, []byte("foo"), []byte("bar"))

	// Test NewIterator
	iter := db.NewIterator(rOption)
	defer iter.Destroy()

	// test iterator SeekToFirst and Next
	for iter.SeekToFirst(); iter.Valid(); iter.Next() {
		if iter.Key() != nil {
			fmt.Println("iter:", string(iter.Key()), string(iter.Value()))
		}
	}

	// test iterator Seek and Prev
	iter.SeekToFirst()
	for iter.Seek([]byte("foo5")); iter.Valid(); iter.Prev() {
		if iter.Key() != nil {
			fmt.Println("iter:", string(iter.Key()), string(iter.Value()))
		}
	}

}
