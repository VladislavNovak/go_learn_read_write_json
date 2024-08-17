package main

import (
	"learn/read_write_json/fileWorker"
	"learn/read_write_json/node"
	"learn/read_write_json/utils"

	"github.com/fatih/color"
)

func createRecord(fileName string) {
	isRepeat := true

	for isRepeat {
		// -- Создаём новое хранилище --
		store, isDone := node.NewStoreExt(fileWorker.NewFileWorker(fileName))

		if !isDone {
			isRepeat = utils.ChooseYesNo("Неудача. Попробовать снова?")
			continue
		}

		// -- Создаём новый узел --
		newNode := node.NewNode()
		// -- Добавляем новый узел --
		store.AddNode(newNode)

		// -- Сохраняем в файл --
		isSave := store.SaveToFile()

		if !isSave {
			isRepeat = utils.ChooseYesNo("Неудача. Попробовать снова?")
			continue
		}

		// Если всё успешно - выходим
		isRepeat = false
	}
}

// Находит записи по url и выводит информацию о них
func printFoundRecords(fileName string) {
	isRepeat := true

	for isRepeat {
		// -- Создаём новое хранилище --
		newStore, isDone := node.NewStoreExt(fileWorker.NewFileWorker(fileName))

		if !isDone {
			isRepeat = utils.ChooseYesNo("Неудача. Попробовать снова?")
			continue
		}

		userInput := utils.GetUserInput("Введите url (либо его часть), чтобы найти записи")
		// -- Получаем коллекцию по условию --
		isCollect := newStore.CollectByUrl(userInput)

		if !isCollect {
			isRepeat = utils.ChooseYesNo("Ничего не найдено. Попробовать снова?")
			continue
		}

		// -- Выводим информацию --
		newStore.Info()

		// Если всё успешно - выходим
		isRepeat = false
	}
}

func deleteRecord(fileName string) {
	isRepeat := true

	for isRepeat {
		// -- Создаём новый узел --
		newStore, isDone := node.NewStoreExt(fileWorker.NewFileWorker(fileName))

		if !isDone {
			isRepeat = utils.ChooseYesNo("Неудача. Попробовать снова?")
			continue
		}

		userInput := utils.GetUserInput("Введите url (либо его часть), чтобы найти записи")
		// -- Получаем коллекцию по условию --
		isCollect := newStore.DeleteByUrl(userInput)

		if !isCollect {
			isRepeat = utils.ChooseYesNo("Ничего не найдено. Попробовать снова?")
			continue
		}

		// -- Выводим информацию --
		newStore.Info()

		// -- Сохраняем выбранную коллекцию (будут удалены записи "к удалению") --
		if utils.ChooseYesNo("Внимание! Сохранить указанную коллекцию?") {
			newStore.SaveToFile()
		}

		// Если всё успешно - выходим
		isRepeat = false
	}
}

func printInfo(fileName string) {
	isRepeat := true

	for isRepeat {
		// -- Создаём новое хранилище --
		newStore, isDone := node.NewStoreExt(fileWorker.NewFileWorker(fileName))

		if !isDone {
			isRepeat = utils.ChooseYesNo("Неудача. Попробовать снова?")
			continue
		}

		// -- Выводим информацию --
		newStore.Info()

		// Если всё успешно - выходим
		isRepeat = false
	}
}

// Создание map из двух массивов ([]string, []func(string))
func createActionList(keys []string, values []func(string)) map[string]func(string) {
	actionList := make(map[string]func(string), len(keys))

	for k, v := range keys {
		actionList[v] = values[k]
	}

	return actionList
}

func main() {
	fileName := "account.json"
	isProcess := true
	menu := [5]string{"Create", "Find", "Remove", "Info", "Exit"}
	listAction := []func(string){createRecord, printFoundRecords, deleteRecord, printInfo}
	actions := createActionList(menu[:len(listAction)], listAction)

	for isProcess {
		selected := utils.SelectFromOptions(menu[:], "Действия с аккаунтом")
		doAction := actions[selected]

		if doAction == nil {
			isProcess = false
			continue
		}

		doAction(fileName)
	}

	color.New(color.FgGreen).Print("Программа завершена")
}
