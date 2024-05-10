package main

// From this
/*
import "fmt"

func myExecuteFn(s string) {
	fmt.Println(s)
}

func main() {
	Execute(myExecuteFn)
}

// this is comming from other party lib
type ExecuteFunc func(string)

func Execute(fn ExecuteFunc) {
	fn("foo bar baz")
}

*/

// to this
import (
	"fmt"
	"net/http"
)

type DB interface {
	Store(string)
}

type Store struct{}

func (s *Store) Store(value string) {
	fmt.Printf("Storing into db %s\n", value)
}

func myExecuteFn(db DB) ExecuteFunc {
	return func(s string) {
		fmt.Println(s)
		db.Store(s)
	}
}

func makeHTTPHandler(db DB, fn httpFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(db, w, r); err != nil {
			return
		}
		db.Store("whatever you want from http server")
	}
}

func main() {
	db := &Store{}
	http.HandleFunc("/", makeHTTPHandler(db, handler))
	Execute(myExecuteFn(db))
}

type httpFunc func(db DB, w http.ResponseWriter, r *http.Request) error

func handler(db DB, w http.ResponseWriter, r *http.Request) error {
	return nil
}

// this is comming from other party lib
type ExecuteFunc func(string)

func Execute(fn ExecuteFunc) {
	fn("foo bar baz")
}
