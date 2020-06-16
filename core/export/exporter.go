package export

import (
	"biu/models/dto"
	"path"
)

// 输出器接口
type Exporter interface {
	Write(info *dto.GenFileInfo) (string, error)
	WriteAsync(ch <-chan *dto.GenFileInfo) <-chan string
}

// 输出器父类
type baseExport struct{}

// 一次性输出
func (exp *baseExport) baseWrite(info *dto.GenFileInfo, expfunc func(filepath, data string)) (string, error) {
	expfunc(path.Join(info.Path, info.Name), info.Content)
	return info.Name, nil
}

// 管道输出
func (exp *baseExport) baseWriteAsync(ch <-chan *dto.GenFileInfo, expfunc func(filepath, data string)) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for info := range ch {
			expfunc(path.Join(info.Path, info.Name), info.Content)
			out <- info.Name
		}
	}()
	return out
}
