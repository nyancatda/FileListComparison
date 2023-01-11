/*
 * @Author: NyanCatda
 * @Date: 2022-11-23 14:11:17
 * @LastEditTime: 2023-01-12 02:50:44
 * @LastEditors: NyanCatda
 * @Description: 参数读取
 * @FilePath: \FileListComparison\internal\Flag\Flag.go
 */
package Flag

import "flag"

type Flag struct {
	SrcPath  string // 源文件夹路径
	DestPath string // 目标文件夹路径
	Copy     bool   // 是否复制出存在差异的文件
}

var Flags Flag // 全局参数变量

/**
 * @description: 初始化参数
 * @return {error} 错误信息
 */
func Init() error {
	// 参数解析
	SrcPath := flag.String("src", "", "源文件夹路径")
	DestPath := flag.String("dest", "", "目标文件夹路径")
	Copy := flag.Bool("copy", false, "是否复制出存在差异的文件")
	flag.Parse()

	// 参数写入变量
	Flags.SrcPath = *SrcPath
	Flags.DestPath = *DestPath
	Flags.Copy = *Copy

	return nil
}
