package fileWorker

import (
	"learn/read_write_json/utils"
	"os"
)

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
