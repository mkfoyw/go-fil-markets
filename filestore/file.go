package filestore

import (
	"os"
	"path"
)

// 文件表示
type fd struct {
	*os.File
	filename string
	basepath string
}

// newFile 创建一个文件
func newFile(basepath OsPath, filename Path) (File, error) {
	var err error
	result := fd{filename: string(filename), basepath: string(basepath)}
	full := path.Join(string(basepath), string(filename))
	result.File, err = os.OpenFile(full, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Path 返回文件相对于存储空间的路径
func (f fd) Path() Path {
	return Path(f.filename)
}

// OsPath 返回文件在文件系统的完整路径
func (f fd) OsPath() OsPath {
	return OsPath(f.Name())
}

// Size 返回文件的大小
func (f fd) Size() int64 {
	info, err := os.Stat(f.Name())
	if err != nil {
		return -1
	}
	return info.Size()
}
