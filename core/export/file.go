package export

import (
	"biu/models/dto"
	"biu/utils/fileUtil"
)

// 文件输出器
type FileExport struct {
	baseExport
}

func (exp *FileExport) ExpFunc(filepath, data string) {
	fileUtil.Write2File(filepath, data)
}

func (exp *FileExport) Write(info *dto.GenFileInfo) (string, error) {
	return exp.baseWrite(info, exp.ExpFunc)
}
func (exp *FileExport) WriteAsync(ch <-chan *dto.GenFileInfo) <-chan string {
	return exp.baseWriteAsync(ch, exp.ExpFunc)
}
