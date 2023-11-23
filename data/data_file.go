package data

import (
	"bitcask-go/fio"
	"fmt"
	"path/filepath"
)

const dataFileSuffix string = ".data"

type DataFile struct {
	Fid         uint32        // 文件ID
	WriteOffset int64         // 写入偏移量
	Manager     fio.IOManager // 文件操作句柄
}

// OpenDataFile 打开数据文件
func OpenDataFile(fileId uint32, filePath string, ioType fio.FileIOType) (*DataFile, error) {
	// 确认文件名，文件名的格式用正则表达式表示是^\d{9}.data，示例000000001.data
	fileName := GetFileName(fileId, filePath)
	// 新建该文件
	return NewDataFile(fileName, fileId, ioType)
}

func NewDataFile(fileName string, fileId uint32, ioType fio.FileIOType) (*DataFile, error) {
	manager, err := fio.NewIOManager(fileName, ioType)
	if err != nil {
		return nil, err
	}
	d := &DataFile{
		Fid: fileId,
		WriteOffset: 0, // 新建文件从0开始
		Manager: manager,
	}
	return d, nil
}

// GetFileName 获取文件名称
func GetFileName(fileId uint32, filePath string) string {
	return filepath.Join(filePath, fmt.Sprintf("%09d%s", fileId, dataFileSuffix))
}

func (file *DataFile) Close() error {
	return nil
}

// Write 将数据写入数据文件中
func (file *DataFile) Write(data []byte) error {
	if n, err := file.Manager.Write(data); err != nil {
		return err
	} else {
		file.WriteOffset += int64(n)
	}
	return nil
}

// Sync 同步文件，利用sync()完成刷盘
func (file *DataFile) Sync() error {
	return file.Manager.Sync()
}

// ReadLogRecord 从偏移量开始读取数据
// TODO 现在参数里的valueSize就是保存在索引里的valueSize，以后需要删除
func (file *DataFile) ReadLogRecord(offset int64) (*LogRecord, error) {
	//TODO value中实际保存的东西除了k-v之外还有CRC等信息，需要提供方法从其中读取出来
	_, err := file.readNBytes(0, offset)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

//读取文件到字节数组，从offset位置开始，读取到长度为n的数组中
func (file *DataFile) readNBytes(n, offset int64) ([]byte, error) {
	b := make([]byte, n)
	_, err := file.Manager.Read(b, offset)
	return b, err
}