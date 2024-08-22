package node

import (
	"encoding/json"
	"learn/read_write_json/encrypter"
	"learn/read_write_json/utils"
	"strings"
	"time"

	"github.com/fatih/color"
)

type iDataBase interface {
	GetFileName() string
	SetFileName(string)
	Write([]byte) bool
	Read() ([]byte, bool)
}

type Store struct {
	Nodes    []Node    `json:"nodes"`
	UpdateAt time.Time `json:"updateAt"`
}

type StoreExt struct {
	Store     Store
	db        iDataBase
	encrypter encrypter.Encripter
}

// (nil, false) - критическая. Файл существует, но невозможно прочитать
// (object, true) - удалось прочитать данные
// (object, false) - файла не было. Создано новое хранилище
func NewStoreExt(dataBase iDataBase, encrypter *encrypter.Encripter) (*StoreExt, bool) {
	storeExt := &StoreExt{
		db:        dataBase,
		encrypter: *encrypter,
	}

	bytes, isRead := storeExt.db.Read()

	// Если удалось прочитать открытую часть
	if isRead {
		var store *Store

		err := json.Unmarshal(bytes, &store)
		if utils.HasError(err, "NewStore/Public/Unmarshal") {
			return nil, false
		}

		if storeExt.encrypter.HasKey() {
			fileName := storeExt.db.GetFileName()
			storeExt.db.SetFileName(fileName + ".encr")

			privateBytes, isPrivateRead := storeExt.db.Read()

			// Если удалось прочитать открытую часть
			if isPrivateRead {
				passwords := make([]string, 0, len(store.Nodes))
				decryptedBytes := storeExt.encrypter.DoDecript(privateBytes)

				if err := json.Unmarshal(decryptedBytes, &passwords); err != nil {
					color.New(color.FgCyan).Println("Ключи дешифровать не удалось!")
				} else {
					for k := range store.Nodes {
						store.Nodes[k].Password = passwords[k]
					}
				}
			}

			defer storeExt.db.SetFileName(fileName)
		}

		storeExt.Store = *store

		return storeExt, true
	}

	// Либо создаём новый стор с пустыми значениями
	// !Предусмотреть случай, при котором, в итоге, значения не должны перетирать старые
	storeExt.Store = Store{}

	return storeExt, false
}

// -- МЕТОДЫ --

// Вернёт массив битов и true, если парсинг прошел успешно
func (s *StoreExt) convertToBytes(data any) ([]byte, bool) {
	bytes, err := json.Marshal(data)
	if utils.HasError(err, "StoreExt/ConvertToBytes") {
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
	// -- Сохраняем приватную часть в отдельный файл --
	if storeExt.encrypter.HasKey() {
		passwords := make([]string, 0, len(storeExt.Store.Nodes))

		for k, v := range storeExt.Store.Nodes {
			passwords = append(passwords, v.Password)
			storeExt.Store.Nodes[k].Password = "ENCRYPTED"
		}

		bytes, isConvert := storeExt.convertToBytes(passwords)
		if !isConvert {
			return false
		}

		encryptedBytes, isEncrypt := storeExt.encrypter.DoEncrypt(bytes)
		if isEncrypt {
			fileName := storeExt.db.GetFileName()
			storeExt.db.SetFileName(fileName + ".encr")
			storeExt.db.Write(encryptedBytes)
			storeExt.db.SetFileName(fileName)
		}
	}

	// -- Сохраняем публичную часть в отдельный файл --
	bytes, isConvert := storeExt.convertToBytes(storeExt.Store)
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
