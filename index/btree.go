package index

import (
	"bitcask-go/data"
	"sync"

	"github.com/google/btree"
)

type BTree struct {
	tree *btree.BTree
	lock *sync.RWMutex
}

func (bt *BTree) Put(key []byte, value *data.LogRecordPos) bool {
	bt.lock.Lock()
	defer bt.lock.Unlock()
	item := Item{key: key, pos: value}
	bt.tree.ReplaceOrInsert(&item)
	return true
}
func (bt *BTree) Get(key []byte) *data.LogRecordPos {
	return nil
}
func (bt *BTree) Delete(key []byte) bool {
	return false
}
