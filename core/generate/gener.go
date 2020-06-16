package generate

import (
	"biu/models/dto"
	"biu/models/tmpl"
)

// 输出器接口
type Gener interface {
	Do(entity *tmpl.GenEntityModel) *dto.GenFileInfo
	DoAsync(ch <-chan *tmpl.GenEntityModel) <-chan *dto.GenFileInfo
}

// 输出器父类
type baseGener struct{}

// 一次性输出
func (exp *baseGener) baseDo(entity *tmpl.GenEntityModel, genfunc func(entity *tmpl.GenEntityModel) *dto.GenFileInfo) *dto.GenFileInfo {
	return genfunc(entity)
}

// 管道输出
func (exp *baseGener) baseDoAsync(ch <-chan *tmpl.GenEntityModel, genfunc func(entity *tmpl.GenEntityModel) *dto.GenFileInfo) <-chan *dto.GenFileInfo {
	out := make(chan *dto.GenFileInfo)
	go func() {
		defer close(out)
		for info := range ch {
			out <- genfunc(info)
		}
	}()
	return out
}
