package generate

import (
	"biu/models/config"
	"biu/models/dto"
	"biu/models/tmpl"
	"bytes"
	"html/template"
	"path"
)

type TmplGen struct {
	baseGener
}

func (tg *TmplGen) DoFunc(entity *tmpl.GenEntityModel) *dto.GenFileInfo {
	res := &dto.GenFileInfo{}
	tmplpath := path.Join(config.Sgy.RootPath, "resources", "tmpls", config.Sgy.Code, "entity.tpl")
	tmpl, err := template.ParseFiles(tmplpath)
	if err != nil {
		panic(err)
	}
	content := bytes.NewBuffer([]byte{})
	err = tmpl.Execute(content, entity)
	if err != nil {
		panic(err)
	}
	res.Content = content.String()
	res.Name = entity.EntityName + "." + config.Sgy.Code
	res.Path = path.Join(config.Sgy.RootPath, "output")
	return res
}

func (tg *TmplGen) Do(entity *tmpl.GenEntityModel) *dto.GenFileInfo {
	return tg.baseDo(entity, tg.DoFunc)
}
func (tg *TmplGen) DoAsync(ch <-chan *tmpl.GenEntityModel) <-chan *dto.GenFileInfo {
	return tg.baseDoAsync(ch, tg.DoFunc)
}
