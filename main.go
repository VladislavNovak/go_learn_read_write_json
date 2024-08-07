package main

import (
	"fmt"
	"learn/read_write_json/account"
	"learn/read_write_json/recorder"
)

func getUserInput(title string) (userInput string) {
	fmt.Printf("Введите %s: ", title)
	fmt.Scan(&userInput)
	return
}

func main() {
	userInputLogin := getUserInput("логин")
	isRepeat := true

	for isRepeat {
		if account, err := account.NewAccount(userInputLogin, ""); err == nil {
			account.PrintData()

			bytes := recorder.ConvertStructToBytes(account)
			recorder.WriteToFile("account.json", bytes)
			isRepeat = false
		}
	}
}
