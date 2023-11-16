package fio

import "os"

type FileIO struct {
	fd *os.File
}

func NewFileIO(fileName string) (*FileIO, error) {
	if f, err := os.OpenFile(fileName,
		os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644); err != nil {
		return nil, err
	} else {
		fio := &FileIO{
			fd: f,
		}
		return fio, nil
	}
}

func (f *FileIO) Read(b []byte, offset int64) (int, error) {
	return f.fd.ReadAt(b, offset)
}
func (f *FileIO) Write(b []byte) (int, error) {
	return f.fd.Write(b)
}
func (f *FileIO) Sync() error {
	return f.fd.Sync()
}
func (f *FileIO) Close() error {
	return f.fd.Close()
}
