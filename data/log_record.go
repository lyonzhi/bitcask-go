package data

type LogRecordType byte

const (
	// 普通的数据
	LogRecordNormal LogRecordType = iota
	// 标记删除的数据
	LogRecordDeleted
)

type LogRecordPos struct {
	// Fid 文件ID
	Fid uint32
	// Offset 偏移量
	Offset int64
	// Size Value的大小
	Size int64
}

type LogRecord struct {
	// 数据的Key
	Key []byte
	// 真实的数据Value
	Value []byte
	// 标记数据的类型是普通的写入还是删除
	Type LogRecordType
}

// EncodeLogRecord 编码
// 返回编码本身和长度
func EncodeLogRecord(*LogRecord) ([]byte, int64) {
	return nil, 0
}
