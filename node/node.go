package node

import (
	"learn/read_write_json/utils"
	"time"

	"github.com/fatih/color"
)

type Node struct {
	Login    string    `json:"login"`
	Password string    `json:"password"`
	Url      string    `json:"url"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}

func (node *Node) PrintData(count int) {
	c := color.New(color.FgGreen)
	c.Printf("%d: [%s] %s %s", count, node.Login, node.Url, node.Password)
	// Работа с форматом
	c.Printf(", (created at %s)\n", node.CreateAt.Local().Format(time.ANSIC))
	// Можно получить данные о тегах модели
	// el, _ := reflect.TypeOf(node).Elem().FieldByName("Login")
	// c.Println(el.Tag)
}

func NewNode() *Node {
	node := &Node{
		Login:    utils.GetLogin(),
		Url:      utils.GetUrl(),
		Password: utils.GetPassword(),
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	return node
}
