package fileWorker

import (
	"learn/read_write_json/utils"
	"os"
)

type FileWorker struct {
	fileName string
}

func NewFileWorker(inFileName string) *FileWorker {
	return &FileWorker{
		fileName: inFileName,
	}
}

// -- МЕТОДЫ --

// return true - если файл создан
func (f *FileWorker) Write(content []byte) bool {
	// Получаем указанный файл
	file, errCaseCreate := os.Create(f.fileName)

	if utils.HasError(errCaseCreate, "WriteToFile/Create") {
		return false
	}

	defer file.Close()

	// Записываем в файл контент
	_, errCaseWrite := file.Write(content)
	return !utils.HasError(errCaseWrite, "WriteToFile/Write")
}

// Читает из файла
func (f *FileWorker) Read() ([]byte, bool) {
	bytes, err := os.ReadFile(f.fileName)
	if utils.HasError(err, "ReadFromFile") {
		return nil, false
	}

	return bytes, true
}
