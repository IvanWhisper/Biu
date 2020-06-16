package pipline

import (
	"biu/core/export"
	"biu/core/generate"
	"biu/core/pump"
	"biu/core/transform"
	"biu/models/config"
)

type Pipline struct {
	Pper pump.Pumper
	Tfer transform.Transformer
	Geer generate.Gener
	Eter export.Exporter
}

func Defalut() *Pipline {
	return &Pipline{
		Pper: &pump.ExcelPump{},
		Tfer: &transform.Transf{},
		Geer: &generate.TmplGen{},
		Eter: &export.FileExport{},
	}
}

func New(p pump.Pumper, t transform.Transformer, g generate.Gener, e export.Exporter) *Pipline {
	return &Pipline{
		Pper: p,
		Tfer: t,
		Geer: g,
		Eter: e,
	}
}

func (pp *Pipline) Test() {
	config.Sgy = &config.Strategy{
		Index:          config.STRATEGY_GO,
		Code:           "go",
		IsUseNickField: true,
		IsAutoRWField:  true,
		RootPath:       "C:\\Users\\ivan\\Documents\\temp",
		InputStr:       "C:\\Users\\ivan\\Documents\\temp\\source.xlsx",
		OutputPath:     "C:\\Users\\ivan\\Documents\\temp\\output",
	}
	pp.RunAsync()
}
