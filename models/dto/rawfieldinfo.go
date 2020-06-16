package dto

// 字段结构
type RawFieldInfo struct {
	PackageName  string `xorm:"varchar(64) notnull 'TABLE_SCHEMA'"`
	StructName   string `xorm:"varchar(64) notnull 'TABLE_NAME'"`
	StructNick   string `xorm:"varchar(2048) notnull 'TABLE_COMMENT'"`
	FieldName    string `xorm:"varchar(64) notnull 'COLUMN_NAME'"`
	FieldNick    string
	FieldComment string `xorm:"varchar(1024) notnull 'COLUMN_COMMENT'"`
	FieldType    string `xorm:"longtext notnull 'DATA_TYPE'"`
	IsCollection bool
	FieldLength  int32 `xorm:"bigint 'CHARACTER_MAXIMUM_LENGTH'"`
	FieldRemark  string
}
