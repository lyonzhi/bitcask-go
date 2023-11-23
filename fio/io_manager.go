package fio

type FileIOType byte

const (
	StandardFileIO = iota
)

type IOManager interface {
	Read([]byte, int64) (int, error)
	Write([]byte) (int, error)
	Sync() error
	Close() error
}

// 为了实现策略模式
var ioTypeMap = make(map[FileIOType]func(string)(*FileIO, error), 0)

func init() {
	ioTypeMap[StandardFileIO] = NewFileIO
}

// NewIOManager 新建一个IOManager，根据ioType的不同，返回的IO类型也不同
func NewIOManager(fileName string, ioType FileIOType) (IOManager, error) {
	// 利用策略模式则优雅许多
	return ioTypeMap[ioType](fileName)
}