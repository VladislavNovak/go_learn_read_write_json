package account

import (
	"learn/read_write_json/utils"
	"time"

	"github.com/fatih/color"
)

type Account struct {
	Login    string    `json:"login"`
	Password string    `json:"password"`
	Url      string    `json:"url"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}

func (acc *Account) PrintData(count int) {
	c := color.New(color.FgGreen)
	c.Printf("%d: [%s] %s %s", count, acc.Login, acc.Url, acc.Password)
	// Работа с форматом
	c.Printf(", (created at %s)\n", acc.CreateAt.Local().Format(time.ANSIC))
	// Можно получить данные о тегах модели
	// el, _ := reflect.TypeOf(acc).Elem().FieldByName("Login")
	// c.Println(el.Tag)
}

func NewAccount() *Account {
	account := &Account{
		Login:    utils.GetLogin(),
		Url:      utils.GetUrl(),
		Password: utils.GetPassword(),
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	return account
}
