/*
 * @Author: NyanCatda
 * @Date: 2023-01-12 02:18:52
 * @LastEditTime: 2023-03-12 17:24:01
 * @LastEditors: NyanCatda
 * @Description: main file
 * @FilePath: \FileListComparison\main.go
 */
package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/nyancatda/FileListComparison/internal/ArrayComparison"
	"github.com/nyancatda/FileListComparison/internal/File"
	"github.com/nyancatda/FileListComparison/internal/Flag"
)

func main() {
	// 初始化参数
	Flag.Init()

	// 读取源文件夹文件列表
	SrcFileList, err := File.GetFilesList(Flag.Flags.SrcPath)
	if err != nil {
		panic(err)
	}
	SrcFiles := make([]string, len(SrcFileList))
	// 提取源文件夹文件列表中的文件名
	for Index, Value := range SrcFileList {
		SrcFiles[Index] = filepath.Clean(strings.Replace(Value, Flag.Flags.SrcPath, "", -1))
	}

	// 读取目标文件夹文件列表
	DestFileList, err := File.GetFilesList(Flag.Flags.DestPath)
	if err != nil {
		panic(err)
	}
	DestFiles := make([]string, len(DestFileList))
	// 提取目标文件夹文件列表中的文件名
	for Index, Value := range DestFileList {
		DestFiles[Index] = filepath.Clean(strings.Replace(Value, Flag.Flags.DestPath, "", -1))
	}

	// 对比文件列表差异
	New, Missing := ArrayComparison.Comparison(SrcFiles, DestFiles)

	// 输出差异
	fmt.Println("新增的文件列表：")
	for _, Value := range New {
		fmt.Println(Value)
	}
	fmt.Println("缺失的文件列表：")
	for _, Value := range Missing {
		fmt.Println(Value)
	}

	// 输出差异数量
	fmt.Printf("新增了%d个文件，缺失了%d个文件\n", len(New), len(Missing))

	// 判断是否需要复制文件
	if Flag.Flags.Copy {
		// 复制新增的文件
		if len(New) != 0 {
			fmt.Println("正在复制新增的文件")
			File.MKDir("./New")
			for _, Value := range New {
				// 创建文件夹
				_, err := File.MKDir("./New/" + filepath.Dir(Value))
				if err != nil {
					panic(err)
				}
				// 复制文件
				_, err = File.Copy(Flag.Flags.DestPath+"/"+Value, "./New/"+Value)
				if err != nil {
					panic(err)
				}
			}
		}
		// 复制缺失的文件
		if len(Missing) != 0 {
			fmt.Println("正在复制缺失的文件")
			File.MKDir("./Missing")
			for _, Value := range Missing {
				// 创建文件夹
				_, err := File.MKDir("./Missing/" + filepath.Dir(Value))
				if err != nil {
					panic(err)
				}
				// 复制文件
				_, err = File.Copy(Flag.Flags.SrcPath+"/"+Value, "./Missing/"+Value)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}
