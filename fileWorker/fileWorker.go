package fileWorker

import (
	"encoding/json"
	"fmt"
	"learn/read_write_json/node"
	"os"
)

func ConvertToBytes(node *node.Node) []byte {
	bytes, _ := json.Marshal(node)
	return bytes
}

func WriteToFile(fileName string, content []byte) {
	if file, err := os.Create(fileName); err == nil {
		if _, err := file.Write(content); err == nil {
			fmt.Println("Файл создан. Контент записан")
		} else {
			fmt.Println("Невозможно записать контент", err)
		}

		defer file.Close()

	} else {
		fmt.Println("Невозможно создать файл", err)
	}
}

func ReadFromFile(fileName string) {
	if bytes, err := os.ReadFile(fileName); err == nil {
		fmt.Println(string(bytes))
	} else {
		fmt.Println("Не удалось прочитать файл")
	}
}