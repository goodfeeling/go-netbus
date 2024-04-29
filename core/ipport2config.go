package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	// Key 建议长度 6-16 个字符，用于身份校验
	KEY = "SrUnI4jrBw1F6Xo"
	//代理端口
	PORT = "6666"
	//开放端口范围，范围（1024~65535）
	ACCESSPORTRANGE = "10000-20000"
	//服务端地址，格式如 45.12.67.98:6666
	SERVERHOST = "45.12.67.98:6666"
	//隧道条数，默认1，范围[1-5]
	TUNNELCOUNT = "1"
	//内网被代理服务地址及访问端口(多个用逗号隔开)，格式如 192.168.1.100:3389:13389
    //内网IP:内网端口:访问端口
	// local-host-mapping = 192.168.1.100:3389:13389
)


func init() {
	fmt.Println("初始化ip配置文件");
	// 定义配置文件的内容
	configContent := fmt.Sprintf(
		"[server]\nkey= %s\nport = %s\naccess-port-range = %s\n\n[client]\nkey = %s\nserver-host = %s\ntunnel-count = %s\n",
		KEY,PORT,ACCESSPORTRANGE,KEY,SERVERHOST,TUNNELCOUNT,
	); 

	
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

	// 将字符串切片合并成一个字符串，以逗号隔开
	if len(lines) > 0 {
		configContent += fmt.Sprintf("local-host-mapping = %s\n"," " + strings.Join(lines[:len(lines)-1], ","))
		configContent += lines[len(lines)-1] 
	}
	configContent += "\n"
	// 将配置内容写入到config.ini文件中
	configFile, err := os.Create("config.ini")
	if err != nil {
		fmt.Println("创建文件失败:", err)
		return
	}
	defer configFile.Close()

	_, err = configFile.WriteString(configContent)
	if err != nil {
		fmt.Println("写入文件失败:", err)
		return
	}

	fmt.Println("配置已成功写入到config.ini文件")
}
