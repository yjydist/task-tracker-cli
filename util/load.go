package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"task-tracker/model"
)

// Load function
// 从 json 文件中加载任务
func Load() {
	// 打开文件, 如果文件不存在则创建
	file, err := os.OpenFile("./tasks.json", os.O_RDWR|os.O_CREATE, 0666)
	// 如果打开文件失败, 则打印错误信息并返回
	if err != nil {
		log.Fatalf("打开或创建文件失败: %v", err)
		return
	}
	// 关闭文件
	defer file.Close()

	// 获取文件信息
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalf("获取文件信息失败: %v", err)
		return
	}

	// 如果文件为空, 则数组为空直接返回
	if fileInfo.Size() == 0 {
		model.Tasks = []model.Task{}
		return
	}
	// 解析 json 数据
	err = json.NewDecoder(file).Decode(&model.Tasks)
	// 如果解析 json 数据失败, 则打印错误信息并返回
	if err != nil {
		log.Fatalf("解析 json 数据失败: %v", err)
		return
	}
	// 打印任务加载成功信息
	fmt.Println("Tasks loaded successfully")
}
