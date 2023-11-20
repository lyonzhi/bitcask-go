package index

import (
	"bitcask-go/data"
	"bytes"

	"github.com/google/btree"
)

type Index interface {
	// Put 保存数据，成功返回true，否则返回false
	Put(key []byte, value *data.LogRecordPos) bool
	// Get 读取key对应的value
	Get(key []byte) *data.LogRecordPos
	// Delete 删除数据
	Delete(key []byte) bool
}

type Item struct {
	key []byte
	pos *data.LogRecordPos
}

func (ai *Item) Less(bi btree.Item) bool {
	return bytes.Compare(ai.key, bi.(*Item).key) == -1
}