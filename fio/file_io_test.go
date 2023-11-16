package fio

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewManager(t *testing.T) {
	manager, err := NewFileIO(filepath.Join("./", "00001.data"))
	assert.Nil(t, err)
	assert.NotNil(t, manager)
}

func TestWrite(t *testing.T) {
	manager, err := NewFileIO(filepath.Join("./", "00001.data"))
	assert.Nil(t, err)
	assert.NotNil(t, manager)

	offset, err := manager.Write([]byte("foo"))
	assert.Nil(t, err)
	assert.Equal(t, 3, offset)
	offset, err = manager.Write([]byte("bar"))
	assert.Nil(t, err)
	assert.Equal(t, 3, offset)
	defer manager.Close()
}

func TestRead(t *testing.T) {
	manager, err := NewFileIO(filepath.Join("./", "00001.data"))
	assert.Nil(t, err)
	assert.NotNil(t, manager)

	bytes := make([]byte, 3)
	_, err = manager.Read(bytes, 3)
	fmt.Println(string(bytes))
}