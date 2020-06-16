package transform

import (
	"biu/models/dto"
	"biu/models/tmpl"
)

type Transformer interface {
	Do(infos []*dto.RawFieldInfo) (*tmpl.GenEntityModel, error)
	DoAsync(ch <-chan []*dto.RawFieldInfo) <-chan *tmpl.GenEntityModel
}
