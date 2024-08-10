package node

import (
	"encoding/json"
	"learn/read_write_json/fileWorker"
	"learn/read_write_json/utils"
	"time"

	"github.com/fatih/color"
)

type Store struct {
	Nodes    []Node    `json:"nodes"`
	UpdateAt time.Time `json:"updateAt"`
	path     string
}

// Вернёт слайс битов и true, если парсинг прошел успешно
func (store *Store) convertToBytes() ([]byte, bool) {
	bytes, err := json.Marshal(store)
	if utils.HasError(err, "Store/ConvertToBytes") {
		return nil, false
	}

	return bytes, true
}

func (store *Store) SaveToFile() bool {
	bytes, isConvert := store.convertToBytes()
	if !isConvert {
		return false
	}

	isWrite := fileWorker.WriteToFile(store.path, bytes)
	return isWrite
}

func (store *Store) AddNode(node *Node) {
	store.Nodes = append(store.Nodes, *node)
	store.UpdateAt = time.Now()
}

// При создании будет сразу получать информацию об указанном файле
// Если файл не обнаружен, будет создаваться пустой store
// false возвращается лишь в случае, если файл обнаружен, но распарсить его не удалось
func NewStore(fileName string) (*Store, bool) {
	bytes, isRead := fileWorker.ReadFromFile(fileName)

	// Если удалось прочитать - сохраняем в новый store
	if isRead {
		var store *Store

		err := json.Unmarshal(bytes, &store)
		color.New(color.FgMagenta).Println("unmarshal:", err)
		if utils.HasError(err, "NewStore/Unmarshal") {
			return nil, false
		}

		store.path = fileName
		return store, true
	}

	color.New(color.FgCyan).Println("Создано новое хранилище!")
	// Либо создаём новый стор
	store := &Store{
		Nodes:    []Node{},
		UpdateAt: time.Now(),
		path:     fileName,
	}

	return store, true
}
