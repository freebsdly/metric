package metric

import (
	"os"
	"path"
)

// 获取指定路径(文件或目录)的磁盘空间使用情况
func TotalSize(filepath string) (tsize int64, err error) {
	var stat os.FileInfo
	stat, err = os.Lstat(filepath)
	if err != nil {
		return
	}
	tsize += stat.Size()

	if !stat.IsDir() {
		return
	}
	fp, err := os.Open(filepath)
	if err != nil {
		return
	}
	names, err := fp.Readdirnames(-1)
	if err != nil {
		return
	}
	for _, v := range names {
		size, err := TotalSize(path.Join(filepath, v))
		if err != nil {
			continue
		}
		tsize += size
	}
	return

}
