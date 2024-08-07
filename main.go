package main

import (
	"fmt"
	pack_account "learn/read_write_json/account"
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
		if account, err := pack_account.NewAccount(userInputLogin, ""); err == nil {
			account.PrintData()
			isRepeat = false
		}
	}
}
