package fileWorker

import (
	"encoding/json"
	"learn/read_write_json/node"
	"learn/read_write_json/utils"
	"os"
)

// Вернёт слайс битов и true, если парсинг прошел успешно
func ConvertToBytes(node *node.Node) ([]byte, bool) {
	bytes, err := json.Marshal(node)
	if utils.HasError(err, "ConvertToBytes") {
		return nil, false
	}

	return bytes, true
}

// Вернёт true, если файл создан
func WriteToFile(fileName string, content []byte) bool {
	// !Получаем указанный файл
	file, errCaseCreate := os.Create(fileName)

	if utils.HasError(errCaseCreate, "WriteToFile/Create") {
		return false
	}

	defer file.Close()

	// !Записываем в указанный файл
	_, errCaseWrite := file.Write(content)
	return !utils.HasError(errCaseWrite, "WriteToFile/Write")
}

func ReadFromFile(fileName string) ([]byte, bool) {
	bytes, err := os.ReadFile(fileName)
	if utils.HasError(err, "ReadFromFile") {
		return nil, false
	}

	return bytes, true
}
