package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func edit() {
	reader := bufio.NewReader(os.Stdin)

	var targetTime, folderPath string
	var parsedTime time.Time
	var err error

	// 获取并验证用户输入
	for {
		fmt.Println("该程序用于批量修改文件的修改时间")
		fmt.Print("请输入目标时间 (yyyymmdd hhmmss): ")
		targetTime, _ = reader.ReadString('\n')
		targetTime = strings.TrimSpace(targetTime)

		// 加载北京时区
		loc, _ := time.LoadLocation("Asia/Shanghai")

		// 使用北京时区解析时间
		parsedTime, err = time.ParseInLocation("20060102 150405", targetTime, loc)
		if err == nil {
			break
		}
		fmt.Println("时间格式错误，请重新输入。")
	}

	fmt.Println("解析的北京时间为：", parsedTime)

	for {
		fmt.Print("请输入目标文件夹的路径: ")
		folderPath, _ = reader.ReadString('\n')
		folderPath = strings.TrimSpace(folderPath)
		if _, err := os.Stat(folderPath); !os.IsNotExist(err) {
			break
		}
		fmt.Println("文件夹路径不存在，请重新输入。")
	}

	// 遍历文件夹并修改文件时间
	filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		return os.Chtimes(path, parsedTime, parsedTime)
	})

	fmt.Println("修改完成")
}

func main() {
	edit()

	// 程序执行完毕，等待用户输入
	fmt.Println("按回车键退出程序：")
	fmt.Scanln() // 等待用户按下回车键
}
