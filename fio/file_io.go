package fio

import "os"

type FileIO struct {
	fd *os.File
}

// NewFileIO 新建文件IO实例
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

// Read 从给定的offset开始读取len(b)长度的数据到b中，返回读取的字节长度
func (f *FileIO) Read(b []byte, offset int64) (int, error) {
	return f.fd.ReadAt(b, offset)
}

// Write 将b的数据写入文件，写入长度是len(b)，返回写入的长度
func (f *FileIO) Write(b []byte) (int, error) {
	return f.fd.Write(b)
}

// Sync 刷盘，这是一个比较重的操作，可以考虑刷盘的时机如何把握
func (f *FileIO) Sync() error {
	return f.fd.Sync()
}

// Close 关闭文件
func (f *FileIO) Close() error {
	return f.fd.Close()
}
