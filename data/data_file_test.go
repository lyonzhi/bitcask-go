package data

import (
	"bitcask-go/fio"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFileName(t *testing.T) {
	name := GetFileName(1, ".")
	fmt.Println(name)
}

func TestNewDataFile(t *testing.T) {
	d, err := NewDataFile("0000001.data", 1, fio.StandardFileIO)
	assert.Nil(t, err)
	fmt.Println(d)
}