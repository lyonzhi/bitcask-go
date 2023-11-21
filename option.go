package bitcaskgo

type Option struct {
	DataFilePath   string //数据文件的路径
	ActiveFileSize int64  //活跃文件大小阈值
	SyncWrites     bool   //是否在每次写入数据后刷盘
}
