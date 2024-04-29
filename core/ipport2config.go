package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func GetIps() String{
	// 获取ipport文件夹下的所有文件
	files, err := ioutil.ReadDir("ipport")
	if err != nil {
		fmt.Println("读取文件夹失败:", err)
		return
	}

	// 用于存储所有行的字符串切片
	var lines []string

	// 遍历文件夹下的所有文件
	for _, file := range files {
		// 忽略文件夹
		if file.IsDir() {
			continue
		}

		// 读取文件内容
		filePath := filepath.Join("ipport", file.Name())
		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Printf("读取文件 %s 失败: %v\n", filePath, err)
			continue
		}

		// 将每行内容添加到字符串切片中
		lines = append(lines, strings.Split(string(data), "\n")...)
	}

	return " " + strings.Join(lines[:len(lines)-1], ",");
}
