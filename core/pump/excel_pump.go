package pump

import (
	"biu/models/config"
	"biu/models/dto"
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
	"strconv"
)

type ExcelPump struct {
}

func (pump *ExcelPump) DoFunc() map[string][]*dto.RawFieldInfo {
	f, err := excelize.OpenFile(config.Sgy.InputStr)
	if err != nil {
		println(err.Error())
		return nil
	}
	groupMap := make(map[string][]*dto.RawFieldInfo)
	s := f.GetSheetName(f.GetActiveSheetIndex())
	rows := f.GetRows(s)
	// 分组
	for index, row := range rows {
		if index == 0 {
			continue
		}

		l, _ := strconv.Atoi(row[8])
		var isFields bool
		if len(row) > 0 && row[9] == "1" {
			isFields = true
		}

		field := &dto.RawFieldInfo{
			PackageName:  s,
			StructName:   row[0],
			StructNick:   row[1],
			FieldName:    row[2],
			FieldComment: row[4],
			FieldNick:    row[3],
			FieldType:    row[5],
			IsCollection: isFields,
			FieldLength:  int32(l),
			FieldRemark:  row[6],
		}
		_, ok := groupMap[field.StructName]
		if !ok {
			groupMap[field.StructName] = make([]*dto.RawFieldInfo, 0)
		}
		groupMap[field.StructName] = append(groupMap[field.StructName], field)
	}
	return groupMap
}

// 一次性输出
func (pump *ExcelPump) Do() map[string][]*dto.RawFieldInfo {
	return pump.DoFunc()
}

// 管道输出
func (pump *ExcelPump) DoAsync() <-chan []*dto.RawFieldInfo {
	out := make(chan []*dto.RawFieldInfo)
	go func() {
		defer close(out)
		for k, v := range pump.DoFunc() {
			log.Println(k)
			out <- v
		}
	}()
	return out
}
