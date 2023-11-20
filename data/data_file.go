package data

import "bitcask-go/fio"

type DataFile struct {
	Fid         uint32        // 文件ID
	WriteOffset int64         // 写入偏移量
	Manager     fio.IOManager // 文件操作句柄
}

// OpenDataFile 打开数据文件
func OpenDataFile(fileId uint32, filePath string, ) (*DataFile, error) {
	return nil, nil
}

// Write 将数据写入数据文件中
func (file *DataFile) Write(data []byte) error {
	if _, err := file.Manager.Write(data); err != nil {
		return err
	}
	return nil
}

func (file *DataFile) Sync() error {
	return file.Manager.Sync()
}