package node

import (
	"time"
)

type Store struct {
	Nodes    []Node    `json:"nodes"`
	UpdateAt time.Time `json:"updateAt"`
}

// При создании будет сразу получать информацию об указанном файле
// Если файл не обнаружен, будет создаваться пустой store
func NewStore(fileName string) *Store {
	// fileWorker
	store := &Store{
		Nodes:    []Node{},
		UpdateAt: time.Now(),
	}

	return store
}
