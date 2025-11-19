package main

import (
	"0g-advanced-dev-course/homework2"
	"fmt"
)

func main() {
	// 1.创建4g大小的文件
	upFilePath := "4gtestfile"
	size := int64(4 * 1024 * 1024 * 1024)
	// size := int64(40 * 1024 * 1024) //测试用
	homework2.CreateNullFile(upFilePath, size)
	fmt.Println("创建文件成功,文件名为：", upFilePath)
	// 2.上传文件
	fragmentSize := int64(400 * 1024 * 1024)
	// fragmentSize := int64(4 * 1024 * 1024)  //测试用
	txHashs, rootHashs, _ := homework2.Upload(upFilePath, fragmentSize)
	fmt.Println("上传文件成功,txHashs:", txHashs, "rootHashs:", rootHashs)
	// 3.下载文件
	downFileName := "4gtestfile_down"
	homework2.Download(rootHashs, downFileName)
	fmt.Println("下载文件成功,文件名为：", downFileName)
}
