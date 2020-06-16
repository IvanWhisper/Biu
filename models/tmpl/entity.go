package tmpl

// 实体模型
type GenEntityModel struct {
	Base
	PackageName    string           // 类、结构
	HasOtherPkg    bool             // 是否有其他包
	ImportPkgNames []string         // 类、结构
	EntityName     string           // 类、结构
	EntityComment  string           // 类、结构注释
	Fields         []*GenFieldModel // 表字段
}

// 字段结构
type GenFieldModel struct {
	FieldName    string // 字段名
	FieldNick    string // 字段名
	FieldType    string // 字段类型
	IsCollection bool   // 是否集合
	IsNeedNick   bool   // 是否启用
	FieldComment string // 字段注释
}
