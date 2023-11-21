package bitcaskgo

import "errors"

var (
	ErrKeyIsEmpty error = errors.New("Key is empty!")
	ErrKeyNotFound error = errors.New("Key not found!")
	ErrDataFileNotFound error = errors.New("Data file not found!")
)