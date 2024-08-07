package pack_account

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/fatih/color"
)

type Account struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (account *Account) generatePassword(size int) {
	symbols := []rune("abcdefgiklmnopqrstuvwyxz123456789!@#$%^&*()_+")
	for i := 0; i < size; i++ {
		account.Password += string(symbols[rand.Intn(len(symbols))])
	}
}

func (account Account) PrintData() {
	c := color.New(color.FgCyan).Add(color.Underline)
	c.Println("---")
	fmt.Println("Логин:", account.Login)
	fmt.Println("Пароль:", account.Password)
	c.Println("---")
}

func NewAccount(login, password string) (*Account, error) {
	if login == "" {
		return nil, errors.New("login_error")
	}

	account := &Account{
		Login:    login,
		Password: password,
	}

	if account.Password == "" {
		account.generatePassword(5)
	}

	return account, nil
}
