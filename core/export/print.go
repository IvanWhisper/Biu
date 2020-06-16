package export

import (
	"biu/models/dto"
	"fmt"
)

// 打印输出器
type PrintExport struct {
	baseExport
}

func (exp *PrintExport) ExpFunc(filepath, data string) {
	fmt.Println(filepath, data)
}

func (exp *PrintExport) Write(info *dto.GenFileInfo) (string, error) {
	return exp.baseWrite(info, exp.ExpFunc)
}

func (exp *PrintExport) WriteAsync(ch <-chan *dto.GenFileInfo) <-chan string {
	return exp.baseWriteAsync(ch, exp.ExpFunc)
}
