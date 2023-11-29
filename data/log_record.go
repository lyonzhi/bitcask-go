package data

import (
	"encoding/binary"
	"hash/crc32"
)

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

type LogRecordHeader struct {
	Crc       uint32
	Type      LogRecordType
	KeySize   uint32
	ValueSize uint32
}

// EncodeLogRecord 编码
// 返回编码本身和长度
/*
根据论文描述，记录应该是这样的（增加了type，确定数据是不是墓碑数据）：
|crc  |type |key size|value size|key |value|
|4字节|1字节 |4字节   |4字节      |不定|不定|
crc: int
type: byte
key size: uint32
value size: uint32
*/
func EncodeLogRecord(logRecord *LogRecord) ([]byte, int64) {
	// 首先编码Header部分，经过计算最大长度就是15个字节
	header := make([]byte, 15)
	// 第5个元素是固定的
	header[4] = byte(logRecord.Type)
	index := 5
	// 保存key size和value size
	index += binary.PutVarint(header[index:], int64(len(logRecord.Key)))
	index += binary.PutVarint(header[index:], int64(len(logRecord.Value)))
	size := index + len(logRecord.Key) + len(logRecord.Value)
	encBytes := make([]byte, size)
	// copy header，此时的index指向value size的最后一位
	copy(encBytes[:index], header[:index])
	// copy key
	copy(encBytes[index:], logRecord.Key)
	// copy value
	copy(encBytes[index+len(logRecord.Key):], logRecord.Value)
	// 计算CRC，注意前4位是预留给CRC的，不要计算进去
	crcValue := crc32.ChecksumIEEE(encBytes[4:])
	binary.LittleEndian.AppendUint32(encBytes[:4], crcValue)

	return encBytes, int64(size)
}
