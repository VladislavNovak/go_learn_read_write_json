package utils

import (
	"fmt"
	"math/rand"
	"net/url"
	"strconv"

	"github.com/fatih/color"
)

func GetLogin() (login string) {
	for {
		login = GetUserInput("Введите логин")

		if len(login) == 0 {
			color.New(color.FgCyan).Println("Неверный логин. Попробуйте снова")
			continue
		}

		return
	}
}

func GetUrl() (userUrl string) {
	for {
		userUrl = GetUserInput("Введите URL")

		_, err := url.ParseRequestURI(userUrl)
		if HasError(err, "NewAccount") {
			color.New(color.FgCyan).Println("Неверный URL. Попробуйте снова")
			continue
		}

		return
	}
}

func GetPassword() (password string) {
	password = GetUserInput("Введите любой пароль. Либо сгенерируйте, введя символ *")

	if password != "*" {
		return
	}

	for {
		sizeRaw := GetUserInput("Введите длину пароля")

		// Конвертируем string to int
		size, err := strconv.Atoi(sizeRaw)
		if HasError(err, "GetPassword") {
			color.New(color.FgCyan).Println("Неверный URL. Попробуйте снова")
			continue
		}

		return string(generatePassword(size))
	}
}

func GetUserInput(title string) (userInput string) {
	color.New(color.FgMagenta).Add(color.Bold).Printf("%s: ", title)
	fmt.Scan(&userInput)
	return
}

func generatePassword(size int) (password string) {
	symbols := []rune("abcdefghiklmnopqrstuvwyxzABCDEFGHIJKLMNOPQRSTUVWYZ123456789!@#$%^&()_+")
	for i := 0; i < size; i++ {
		password += string(symbols[rand.Intn(len(symbols))])
	}
	return
}

func HasError(err error, srcName string) bool {
	if err != nil {
		color.New(color.FgBlue).Printf("Ошибка %s (функция %s)", err.Error(), srcName)
		return true
	}

	return false
}
