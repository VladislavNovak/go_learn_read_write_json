package node

import (
	"encoding/json"
	"learn/read_write_json/utils"
	"strings"
	"time"

	"github.com/fatih/color"
)

type iDataBase interface {
	Write([]byte) bool
	Read() ([]byte, bool)
}

type Store struct {
	Nodes    []Node    `json:"nodes"`
	UpdateAt time.Time `json:"updateAt"`
}

type StoreExt struct {
	Store Store
	db    iDataBase
}

// false возвращается лишь в случае, если файл обнаружен, но распарсить его не удалось
func NewStoreExt(dataBase iDataBase) (*StoreExt, bool) {
	storeExt := &StoreExt{db: dataBase}

	bytes, isRead := storeExt.db.Read()

	// Если удалось прочитать - сохраняем в store
	if isRead {
		var store *Store

		err := json.Unmarshal(bytes, &store)
		if utils.HasError(err, "NewStore/Unmarshal") {
			return nil, false
		}

		storeExt.Store = *store

		return storeExt, true
	}

	color.New(color.FgCyan).Println("Создано новое хранилище!")

	// Либо создаём новый стор с пустыми значениями
	storeExt.Store = Store{}

	return storeExt, true
}

// -- МЕТОДЫ --

// Конвертирует лишь сущность Store (!не StoreExt)
// Вернёт массив битов и true, если парсинг прошел успешно
func (store *Store) convertToBytes() ([]byte, bool) {
	bytes, err := json.Marshal(store)
	if utils.HasError(err, "Store/ConvertToBytes") {
		return nil, false
	}

	return bytes, true
}

// Предикат. Проверяет вхождение подстроки в login
func checkSubLogin(node Node, subStr string) bool {
	return strings.Contains(node.Login, subStr)
}

// Предикат. Проверяет вхождение подстроки в ull
func checkSubUrl(node Node, subStr string) bool {
	return strings.Contains(node.Url, subStr)
}

func (storeExt *StoreExt) DoDeleteByLogin(subStr string) bool {
	return storeExt.Store.filter("delete", subStr, checkSubLogin)
}

func (storeExt *StoreExt) DoDeleteByUrl(subStr string) bool {
	return storeExt.Store.filter("delete", subStr, checkSubUrl)
}

func (storeExt *StoreExt) DoCollectByLogin(subStr string) bool {
	return storeExt.Store.filter("collect", subStr, checkSubLogin)
}

func (storeExt *StoreExt) DoCollectByUrl(subStr string) bool {
	return storeExt.Store.filter("collect", subStr, checkSubUrl)
}

// mode ("delete" | "collect");
// subStr - критерий фильтрации;
// checkContains - предикат;
// return bool (true, если удалось получить хотя бы один элемент коллекции)
func (store *Store) filter(mode string, subStr string, checkContains func(Node, string) bool) bool {
	if len(store.Nodes) == 0 {
		color.New(color.FgBlue).Printf("store/%s. Данных пока нет\n", mode)
		return false
	}

	collection := []Node{}

	switch mode {
	case "collect":
		// -- Соберём временную коллекцию --
		for _, v := range store.Nodes {
			if checkContains(v, subStr) {
				collection = append(collection, v)
			}
		}
	case "delete":
		for k, v := range store.Nodes {
			if !checkContains(v, subStr) {
				collection = append(collection, v)
			} else {
				color.New(color.FgMagenta).Printf("К удалению: ")
				v.PrintData(k)
			}
		}
	}

	if len(collection) == 0 {
		color.New(color.FgBlue).Printf("store/%s. Данных пока нет\n", mode)
		return false
	}

	// -- Передадим временную коллекцию в store --
	store.Nodes = collection
	return true
}

func (storeExt *StoreExt) Info() {
	if len(storeExt.Store.Nodes) == 0 {
		color.New(color.FgBlue).Println("store/Info. Данных пока нет")
		return
	}

	color.New(color.FgGreen).Println("Данные в коллекции:")
	for k, v := range storeExt.Store.Nodes {
		v.PrintData(k)
	}
}

func (storeExt *StoreExt) SaveToFile() bool {
	bytes, isConvert := storeExt.Store.convertToBytes()
	if !isConvert {
		return false
	}

	isWrite := storeExt.db.Write(bytes)
	return isWrite
}

func (storeExt *StoreExt) AddNode(node *Node) {
	storeExt.Store.Nodes = append(storeExt.Store.Nodes, *node)
	storeExt.Store.UpdateAt = time.Now()
}
