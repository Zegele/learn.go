package main

import (
	"fmt"
	"sync"
)

var once sync.Once = sync.Once{} //再看看这个sync.Once{}怎么用。

var facStore = &dbFactoryStore{}

type dbFactoryStore struct {
	store map[string]DBFactory
}

type Conn struct{}

type DBFactory interface {
	GetConnection() *Conn
}

func initMySqlFac(connStr string) DBFactory {
	return &MySqlDBFactory{}
}

type MySqlDBFactory struct {
	once sync.Once
}

func (MySqlDBFactory) GetConnection() *Conn {
	once.Do(func() {
		initMySqlFac("546")
	})
	fmt.Println(facStore)
	// todo
	return nil
}

func main() {

	connStr := "xxxxxx"
	fmt.Println(connStr)
}
