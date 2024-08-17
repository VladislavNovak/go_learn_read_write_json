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
func findRecords(fileName string, mode string, propose string) {
	// func findRecords(fileName string) {
	isRepeat := true

	for isRepeat {
		// -- Создаём новое хранилище --
		newStore, isDone := node.NewStoreExt(fileWorker.NewFileWorker(fileName))

		if !isDone {
			isRepeat = utils.ChooseYesNo("Неудача. Попробовать снова?")
			continue
		}

		var isCollect bool
		// -- Получаем коллекцию по условию --
		switch mode {
		case "url":
			isCollect = newStore.DoCollectByUrl(utils.GetUserInput(propose))
		case "login":
			isCollect = newStore.DoCollectByLogin(utils.GetUserInput(propose))
		}

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

func findRecordsByUrl(fileName string) {
	findRecords(fileName, "url", "Введите url (либо его часть), чтобы найти записи")
}

func findRecordsByLogin(fileName string) {
	findRecords(fileName, "login", "Введите login (либо его часть), чтобы найти записи")
}

func deleteRecords(fileName string, mode string, propose string) {
	isRepeat := true

	for isRepeat {
		// -- Создаём новый узел --
		newStore, isDone := node.NewStoreExt(fileWorker.NewFileWorker(fileName))

		if !isDone {
			isRepeat = utils.ChooseYesNo("Неудача. Попробовать снова?")
			continue
		}

		var isCollect bool
		// -- Получаем коллекцию по условию --
		switch mode {
		case "url":
			isCollect = newStore.DoDeleteByUrl(utils.GetUserInput(propose))
		case "login":
			isCollect = newStore.DoDeleteByLogin(utils.GetUserInput(propose))
		}

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

func deleteRecordsByUrl(fileName string) {
	deleteRecords(fileName, "url", "Введите url (либо его часть), чтобы найти записи")
}

func deleteRecordsByLogin(fileName string) {
	deleteRecords(fileName, "login", "Введите login (либо его часть), чтобы найти записи")
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
	menu := [7]string{"Create", "Find by URL", "Find by Login", "Remove by URL", "Remove by Login", "Info", "Exit"}
	listAction := []func(string){
		createRecord,
		findRecordsByUrl,
		findRecordsByLogin,
		deleteRecordsByUrl,
		deleteRecordsByLogin,
		printInfo,
	}

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
