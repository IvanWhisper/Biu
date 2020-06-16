package pump

import (
	"biu/models/dto"
)

type Pumper interface {
	Do() map[string][]*dto.RawFieldInfo
	DoAsync() <-chan []*dto.RawFieldInfo
}
