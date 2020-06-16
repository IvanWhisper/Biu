//*******************************
// Create By {{.Author}}
// Date {{.CreateTime}}
// Code Generate By biu
//*******************************

using System;
using Newtonsoft.Json;
{{ if .HasOtherPkg }}{{range $i, $pkgname := .ImportPkgNames}}using {{$pkgname}};
{{end}}{{ else }}{{ end }}
namespace {{.PackageName}}
{
    ///<summary>
    /// {{.EntityComment}}
    ///</summary>
    public class {{.EntityName}}
    {
{{range $j, $item := .Fields}}
        ///<summary>
        /// {{$item.FieldComment}}
        ///</summary>{{if .IsNeedNick}}
        [JsonProperty("{{.FieldName}}")]{{else}}{{end}}
        public {{if .IsCollection}}{{$item.FieldType}}[]{{else}}{{$item.FieldType}}{{end}} {{$item.FieldNick}} { get; set; }
{{end}}
     }
}





