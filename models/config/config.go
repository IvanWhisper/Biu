package config

type StrategyType int32

const (
	STRATEGY_GO StrategyType = 0
	STRATEGY_CS StrategyType = 1
)

var Sgy *Strategy

type Strategy struct {
	Index          StrategyType
	Code           string
	IsUseNickField bool
	IsAutoRWField  bool
	RootPath       string
	InputStr       string
	OutputPath     string
}

//func New(st StrategyType)Strategy{
//	if st==STRATEGY_GO{
//		return Strategy{STRATEGY_GO,"go",true,true,fileUtil.GetExeRootDir()}
//	}else{
//		return Strategy{STRATEGY_CS,"cs",true,true,fileUtil.GetExeRootDir()}
//	}
//}
