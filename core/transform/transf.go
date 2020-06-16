package transform

import (
	"biu/models/config"
	"biu/models/dto"
	"biu/models/tmpl"
	"log"
)

type Transf struct {
}

// 一次性输出
func (tf *Transf) Do(infos []*dto.RawFieldInfo) (*tmpl.GenEntityModel, error) {
	return Fields2Entity(infos, config.Sgy)
}

// 管道输出
func (tf *Transf) DoAsync(ch <-chan []*dto.RawFieldInfo) <-chan *tmpl.GenEntityModel {
	out := make(chan *tmpl.GenEntityModel)
	go func() {
		defer close(out)
		for info := range ch {
			res, e := Fields2Entity(info, config.Sgy)
			if e != nil {
				log.Println(e)
			} else {
				out <- res
			}
		}
	}()
	return out
}
