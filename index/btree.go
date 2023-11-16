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

func NewBtree() *BTree {
	return &BTree{
		tree: btree.New(32),
		lock: new(sync.RWMutex),
	}
}

func (bt *BTree) Put(key []byte, value *data.LogRecordPos) bool {
	bt.lock.Lock()
	defer bt.lock.Unlock()
	item := &Item{key: key, pos: value}
	bt.tree.ReplaceOrInsert(item)
	return true
}
func (bt *BTree) Get(key []byte) *data.LogRecordPos {
	k := &Item{key: key}
	item := bt.tree.Get(k)
	if item == nil {
		return nil
	}
	return item.(*Item).pos
}
func (bt *BTree) Delete(key []byte) bool {
	bt.lock.Lock()
	defer bt.lock.Unlock()
	k := &Item{key: key}
	item := bt.tree.Delete(k)
	return item != nil
}
