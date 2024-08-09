package main

import (
	"learn/read_write_json/fileWorker"
	"learn/read_write_json/node"
	"learn/read_write_json/utils"
	"strings"
)

func main() {
	fileName := "account.json"

	for {
		node := node.NewNode()
		node.PrintData(0)

		bytes, isConvert := fileWorker.ConvertToBytes(node)
		if !isConvert {
			continue
		}

		isWrite := fileWorker.WriteToFile(fileName, bytes)
		if !isWrite {
			continue
		}

		userInput := utils.GetUserInput("Продолжить (Y)?")
		if strings.Contains("Yy", userInput) {
			continue
		}

		break
	}
}
