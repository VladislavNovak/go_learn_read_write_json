package main

import (
	"learn/read_write_json/fileWorker"
	"learn/read_write_json/node"
	"learn/read_write_json/utils"
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

func main() {
	menu := [5]string{"Create", "Find", "Remove", "Info", "Exit"}
	fileName := "account.json"
	isProcess := true

	for isProcess {
		selected := utils.SelectFromOptions(menu[:], "Действия с аккаунтом")
		switch selected {
		case "Create":
			createRecord(fileName)
		case "Find":
			printFoundRecords(fileName)
		case "Remove":
			deleteRecord(fileName)
		case "Info":
			printInfo(fileName)
		default:
			isProcess = false
		}
	}
}
