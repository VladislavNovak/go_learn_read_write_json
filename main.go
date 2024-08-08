package main

import (
	"learn/read_write_json/account"
	"learn/read_write_json/recorder"
)

func main() {
	fileName := "account.json"

	isRepeat := true

	for isRepeat {
		account := account.NewAccount()
		account.PrintData(0)

		bytes := recorder.ConvertToBytes(account)
		recorder.WriteToFile(fileName, bytes)
		isRepeat = false
	}
}
