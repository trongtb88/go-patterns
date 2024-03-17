package main

import "fmt"

type DB interface {
	Store(s string) error
}

type Store struct {
}

func (s *Store) Store(value string) error {
	fmt.Println("Storing into db ", value)
	return nil
}

func myExecuteFunc(db DB) ExecuteFunc {
	return func(s string) {
		fmt.Println(s)
		db.Store(s)
	}
}

type ExecuteFunc func(string)

func Execute(fn ExecuteFunc) {
	fn("abc")
}

func main() {
	db := &Store{}
	Execute(myExecuteFunc(db))
}
