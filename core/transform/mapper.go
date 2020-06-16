package transform

import (
	"biu/models/config"
	"biu/models/dto"
	"biu/models/tmpl"
	"biu/utils/fileUtil"
	"errors"
	"log"
	"strings"
	"time"
)

// 对象转化
func Fields2Entity(infos []*dto.RawFieldInfo, sgy *config.Strategy) (*tmpl.GenEntityModel, error) {
	if len(infos) <= 0 {
		log.Println("没有字段")
		return nil, errors.New("空")
	} else {
		result := &tmpl.GenEntityModel{}
		result.PackageName = infos[0].PackageName

		if len(infos[0].StructNick) > 0 {
			result.EntityName = infos[0].StructNick
		} else {
			result.EntityName = EntityNameFormat(infos[0].StructName)
		}

		result.EntityComment = infos[0].StructName

		result.Author = "Ivan"
		result.CreateTime = time.Now().Format("2006-01-02 15:04:05")

		result.HasOtherPkg = false

		fields := make([]*tmpl.GenFieldModel, 0)
		for _, info := range infos {
			field := &tmpl.GenFieldModel{}
			field.FieldType = FieldTypeFormat(info.FieldType, sgy)
			field.FieldName = info.FieldName
			field.FieldNick = FieldNickFormat(info.FieldNick, info.FieldName, field.FieldType, sgy)
			field.IsNeedNick = !(field.FieldName == field.FieldNick)
			field.FieldComment = info.FieldComment + "(" + info.FieldRemark + ")"
			field.IsCollection = info.IsCollection
			fields = append(fields, field)
		}
		result.Fields = fields
		return result, nil
	}
}

func FieldNickFormat(nick, name, fieldtype string, sgy *config.Strategy) string {
	var res string
	if sgy.IsUseNickField && len(nick) > 0 {
		res = nick
	} else if sgy.IsAutoRWField {
		res = fileUtil.Capitalize(name)
	} else {
		res = name
	}
	if res == fieldtype {
		res += "Value"
	}
	return res
}

func EntityNameFormat(name string) string {
	return fileUtil.Capitalize(name)
}

func FieldTypeFormat(fieldtype string, sgy *config.Strategy) string {
	curType := strings.ToLower(fieldtype)
	if v, ok := TypeMap[sgy.Index][curType]; ok {
		return v
	} else {
		return EntityNameFormat(fieldtype)
	}
}
