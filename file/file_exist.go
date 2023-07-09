package file

import "os"

/*
FileExist
判断文件是否存在
Param:
	path string 文件的位置
Return:
	bool 如果存在则为真
*/

func FileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
