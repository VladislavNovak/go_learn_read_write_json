package main

import (
	"learn/read_write_json/node"
	"learn/read_write_json/utils"
)

func createAccount(fileName string) {
	isRepeat := true

	for isRepeat {
		newNode := node.NewNode()
		store, isDone := node.NewStore(fileName)

		if !isDone {
			isRepeat = utils.ChooseYesNo("Неудача. Попробовать снова?")
			continue
		}

		// Добавляем новый узел
		store.AddNode(newNode)

		// Сохраняем в файл
		isSave := store.SaveToFile()

		if !isSave {
			isRepeat = utils.ChooseYesNo("Неудача. Попробовать снова?")
			continue
		}

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
			createAccount(fileName)
		default:
			isProcess = false
		}
	}
}
