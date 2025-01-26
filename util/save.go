package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"task-tracker/model"
)

// Save function
// 保存任务列表到 json 文件
func Save() {
	// 以覆盖模式打开文件（写入模式 + 自动创建文件 + 清空文件内容）
	file, err := os.OpenFile("./tasks.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("Save: 打开文件失败: %v", err)
		return
	}
	defer file.Close()

	// 将结构体数组编码为 JSON 数据
	tasksJson, err := json.Marshal(model.Tasks)
	if err != nil {
		log.Fatalf("Save: 编码 JSON 数据失败: %v", err)
		return
	}

	// 写入 JSON 数据到文件
	_, err = file.Write(tasksJson)
	if err != nil {
		log.Fatalf("Save: 写入 JSON 数据失败: %v", err)
		return
	}

	fmt.Println("Tasks saved successfully")
}
