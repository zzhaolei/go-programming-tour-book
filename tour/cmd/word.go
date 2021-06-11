// Package cmd
// 在Golang中 一个包中数据的加载顺序为
// 1。按照顺序执行import的包
// 2。加载const
// 3。加载var
// 4。加载init，如果存在多个init，则按照字典序执行
package cmd

import (
	"log"
	"strings"

	"github.com/zzhaolei/go-programming-tour-book/tour/internal/word"

	"github.com/spf13/cobra"
)

const (
	ModeUpper                      = iota + 1 // 单词转为大写
	ModeLower                                 // 单词转为小写
	ModeUnderscoreToUpperCamelCase            // 下划线单词转为大写
	ModeUnderscoreToLowerCamelCase            // 下划线单词转为小写
	ModeCamelCaseToUnderscore                 // 驼峰单词转为下划线单词
)

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下：",
	"    1：单词转为大写",
	"    2：单词转为小写",
	"    3：下划线单词转为大写",
	"    4：下划线单词转为小写",
	"    5：驼峰单词转为下划线单词",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatal("暂不支持该转换模式，请执行 help word 查看帮助文档")
		}

		log.Printf("转换结果: \n%s\n", content)
	},
}

var str string
var mode int8

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的格式")
}
