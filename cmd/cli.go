package cmd

import (
	"biu/core/pipline"
	"biu/models/config"
	"biu/utils/fileUtil"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/urfave/cli"
	"os"
	"path"
)

func Run() error {
	app := cli.NewApp()
	app.Name = "Biu"
	app.Usage = "简单对象生成器"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "operation,o",
			Value: "gen",
			Usage: "operation",
		},
		cli.StringFlag{
			Name:  "type, t",
			Value: "excel",
			Usage: "source type",
		},
		cli.StringFlag{
			Name:  "data, d",
			Value: "./",
			Usage: "KeyString",
		},
		cli.StringFlag{
			Name:  "lang, l",
			Value: "CS",
			Usage: "language",
		},
	}
	app.Action = func(c *cli.Context) error {
		operation := c.String("operation")
		sourcetype := c.String("type")
		data := c.String("source")
		lang := c.String("lang")

		var sourcePath string
		if operation == "gen" {
			// excel 为源
			if sourcetype == "excel" {
				sourcePath = path.Join(data, "source.xlsx")
				sIsEx := fileUtil.IsDirOrFileExist(sourcePath)
				if !sIsEx {
					panic("源文件不存在")
				}
			}

			// 输出路径
			outputPath := path.Join(data, "output")
			opIsEx := fileUtil.IsDirOrFileExist(outputPath)
			if !opIsEx {
				err := os.MkdirAll(outputPath, os.ModePerm)
				if err != nil {
					fmt.Println(err)
				}
			}
			// 构建语言类型
			var st config.StrategyType
			var code string
			if lang == "CS" {
				st = config.STRATEGY_CS
				code = "cs"
				// 配置
				config.Sgy = &config.Strategy{
					Index:          st,
					Code:           code,
					IsUseNickField: true,
					IsAutoRWField:  false,
					RootPath:       data,
					InputStr:       sourcePath,
					OutputPath:     outputPath,
				}
			} else {
				st = config.STRATEGY_GO
				code = "go"
				// 配置
				config.Sgy = &config.Strategy{
					Index:          st,
					Code:           code,
					IsUseNickField: true,
					IsAutoRWField:  true,
					RootPath:       data,
					InputStr:       sourcePath,
					OutputPath:     outputPath,
				}
			}
			// 排布流水线
			pl := pipline.Defalut()
			// 执行流水线
			pl.RunAsync()
		} else if operation == "c" {
			f := excelize.NewFile()
			// Create a new sheet.
			heads := map[string]string{
				"A1": "EntityName(类名)",
				"B1": "EntityNick(别名)",
				"C1": "FieldName(字段)",
				"D1": "FieldNick(字段名)",
				"E1": "FieldComment(字段注释)",
				"F1": "FieldType(字段类型)",
				"G1": "FieldRemark(备注)",
				"H1": "Require(是否必备)",
				"I1": "FieldLength(字段长度)",
				"J1": "IsList(是否集合，1-表示是)",

				"A2": "wdt_res",
				"B2": "WdtRes",
				"C2": "code",
				"D2": "WdtCode",
				"E2": "错误码",
				"F2": "int",
				"G2": "状态码:0表示成功,其他表示失败",
				"H2": "是",
				"I2": "11"}
			// Set value of a cell.
			for k, v := range heads {
				f.SetCellValue("sheet1", k, v)
			}
			// Save xlsx file by the given path.
			if err := f.SaveAs(path.Join(data, "source.xlsx")); err != nil {
				fmt.Println(err)
			}
		}
		return nil
	}
	return app.Run(os.Args)
}
