package node

import (
	"encoding/json"
	"learn/read_write_json/fileWorker"
	"learn/read_write_json/utils"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Store struct {
	Nodes    []Node    `json:"nodes"`
	UpdateAt time.Time `json:"updateAt"`
	path     string
}

// Вернёт true, если удалось получить хотя бы один элемент коллекции
func (store *Store) filterByUrl(url string, action string) bool {
	if len(store.Nodes) == 0 {
		color.New(color.FgBlue).Printf("store/%s. Данных пока нет\n", action)
		return false
	}

	collection := []Node{}

	switch action {
	case "collect":
		// -- Соберём временную коллекцию --
		for _, v := range store.Nodes {
			if strings.Contains(v.Url, url) {
				collection = append(collection, v)
			}
		}
	case "delete":
		for k, v := range store.Nodes {
			if !strings.Contains(v.Url, url) {
				collection = append(collection, v)
			} else {
				color.New(color.FgMagenta).Printf("К удалению: ")
				v.PrintData(k)
			}
		}
	}

	if len(collection) == 0 {
		color.New(color.FgBlue).Printf("store/%s. Данных пока нет\n", action)
		return false
	}

	// -- Передадим временную коллекцию в store --
	store.Nodes = collection
	return true
}

// Вернёт true, если удалось получить хотя бы один элемент коллекции
func (store *Store) DeleteByUrl(url string) bool {
	return store.filterByUrl(url, "delete")
}

// Вернёт true, если удалось получить хотя бы один элемент коллекции
func (store *Store) CollectByUrl(url string) bool {
	return store.filterByUrl(url, "collect")
}

func (store *Store) Info() {
	if len(store.Nodes) == 0 {
		color.New(color.FgBlue).Println("store/Info. Данных пока нет")
		return
	}

	color.New(color.FgGreen).Println("Данные в коллекции:")
	for k, v := range store.Nodes {
		v.PrintData(k)
	}
}

// Вернёт массив битов и true, если парсинг прошел успешно
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
