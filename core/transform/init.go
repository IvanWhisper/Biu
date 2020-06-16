package transform

import "biu/models/config"

var TypeMap = map[config.StrategyType]map[string]string{
	config.STRATEGY_GO: Mysql2GoMap,
	config.STRATEGY_CS: Mysql2CSMap,
}

var Mysql2GoMap = map[string]string{
	"tinyint":    "int32",
	"smallint":   "int32",
	"mediumint":  "int32",
	"int":        "int32",
	"integer":    "int64",
	"bigint":     "int64",
	"float":      "float64",
	"double":     "float64",
	"decimal":    "float64",
	"date":       "string",
	"time":       "string",
	"year":       "string",
	"datetime":   "time.Time",
	"timestamp":  "time.Time",
	"string":     "string",
	"char":       "string",
	"varchar":    "string",
	"tinyblob":   "string",
	"tinytext":   "string",
	"blob":       "string",
	"text":       "string",
	"mediumblob": "string",
	"mediumtext": "string",
	"longblob":   "string",
	"longtext":   "string",
}

var Mysql2CSMap = map[string]string{
	"tinyint":    "Int32",
	"smallint":   "Int32",
	"mediumint":  "Int32",
	"int":        "Int32",
	"integer":    "Int64",
	"bigint":     "Int64",
	"float":      "float",
	"double":     "double",
	"decimal":    "decimal",
	"date":       "string",
	"time":       "string",
	"year":       "string",
	"datetime":   "DateTime",
	"timestamp":  "string",
	"string":     "string",
	"char":       "string",
	"varchar":    "string",
	"tinyblob":   "string",
	"tinytext":   "string",
	"blob":       "string",
	"text":       "string",
	"mediumblob": "string",
	"mediumtext": "string",
	"longblob":   "string",
	"longtext":   "string",
}
