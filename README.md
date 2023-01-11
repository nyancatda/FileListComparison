# FileListComparison
对比两个文件夹内增加或缺失的文件

## ⚙️ 参数
```
-copy
    是否复制出存在差异的文件
-dest string
    目标文件夹路径
-src string
    源文件夹路径
```

## 🎬 使用
以Windows系统为例
```
.\FileListComparison.exe -src ./src_folder -dest ./dest_folder -copy
```
执行后会对比src_folder文件夹与dest_folder文件夹内的差异，并复制出差异文件到程序同目录下的`New`和`Missing`文件夹内

## 📖 许可证
项目采用`Mozilla Public License Version 2.0`协议开源

二次修改源代码需要开源修改后的代码，对源代码修改之处需要提供说明文档