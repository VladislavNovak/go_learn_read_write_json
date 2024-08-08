package main

import (
	"learn/read_write_json/fileWorker"
	"learn/read_write_json/node"
)

func main() {
	fileName := "account.json"

	isRepeat := true

	for isRepeat {
		node := node.NewNode()
		node.PrintData(0)

		bytes := fileWorker.ConvertToBytes(node)
		fileWorker.WriteToFile(fileName, bytes)
		isRepeat = false
	}
}
