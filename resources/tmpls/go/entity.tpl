//*******************************
// Create By {{.Author}}
// Date {{.CreateTime}}
// Code Generate By biu
//*******************************

package {{.PackageName}}
{{ if .HasOtherPkg }}
import (
{{range $i, $pkgname := .ImportPkgNames}}    "{{$pkgname}}"
{{end}})
{{ else }}{{ end }}
// {{.EntityComment}}
type {{.EntityName}} struct {
{{range $j, $item := .Fields}}    {{$item.FieldNick}} {{if .IsCollection}}[]{{$item.FieldType}}{{else}}{{$item.FieldType}}{{end}} `json:"{{.FieldName}}"`// {{$item.FieldComment}}
{{end}}}