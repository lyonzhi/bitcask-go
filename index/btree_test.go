package index

import (
	"bitcask-go/data"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPut(t *testing.T) {
	bt := NewBtree()
	key := "foo"
	v := &data.LogRecordPos{
		Fid: 123,
		Offset: 11111,
		Size: 12,
	}
	res := bt.Put([]byte(key), v)
	assert.True(t, res)
	
}