package bitcaskgo

import (
	"bitcask-go/data"
	"bitcask-go/index"
	"sync"
)

type DB struct {
	mu         *sync.RWMutex
	activeFile *data.DataFile            //活跃数据文件，有且只有一个
	oldFiles   map[uint32]*data.DataFile // 不活跃数据文件，用map保存，key是文件ID
	option     Option
	index      index.Index
}

func (db *DB) Put(key []byte, value []byte) error {
	if len(key) == 0 {
		return ErrKeyIsEmpty
	}
	// 构造Data
	data := &data.LogRecord{
		Key:   key,
		Value: value,
		Type:  data.LogRecordNormal,
	}

	// 将data顺序写入到数据文件中
	logRecordPos, err := db.appendLogRecord(data)
	if err != nil {
		return err
	}

	// 更新数据索引
	db.index.Put(key, logRecordPos) 
	return nil
}

func (db *DB) appendLogRecord(logRecord *data.LogRecord) (*data.LogRecordPos, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	// 判断活跃文件是否存在，如果不存在，新建一个
	if db.activeFile == nil {
		if err := db.setActiveFile(); err != nil {
			return nil, err
		}
	}
	// 写入数据
	code, size := data.EncodeLogRecord(logRecord)
	// 如果写入的数据超过了active file的阈值，就要切换文件，将现在的文件置为old，然后重启一个新的active
	// 阈值由参数决定
	if db.activeFile.WriteOffset + size > db.option.ActiveFileSize {
		if err := db.activeFile.Sync(); err != nil {
			return nil, err
		}
		db.oldFiles[db.activeFile.Fid] = db.activeFile
		if err := db.setActiveFile(); err != nil {
			return nil, err
		}
	}

	if err := db.activeFile.Write(code); err != nil {
		return nil, err
	}

	if db.option.SyncWrites {
		if err := db.activeFile.Sync(); err != nil {
			return nil, err
		}
	}

	// 更新索引
	logRecordPos := &data.LogRecordPos{
		Fid:    db.activeFile.Fid,
		Offset: db.activeFile.WriteOffset,
		Size:   size,
	}
	return logRecordPos, nil
}

// 设置当前活跃文件
func (db *DB) setActiveFile() error {
	var initialFileId uint32 = 0
	if db.activeFile != nil {
		initialFileId = db.activeFile.Fid + 1
	}
	// 打开文件
	if dataFile, err := data.OpenDataFile(initialFileId, db.option.DataFilePath); err != nil {
		return err
	} else {
		db.activeFile = dataFile
	}
	return nil
}
