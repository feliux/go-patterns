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

// DB is the interface for storing data into the database.
type DB interface {
	Store(string)
}

// Store may contain the database configuration.
type Store struct{}

// Store just stores into the databse.
func (s *Store) Store(value string) {
	fmt.Printf("Storing into db %s\n", value)
}

// myEcexuteFn executes a custom function with the ExecuteFunc signature.
func myExecuteFn(db DB) ExecuteFunc {
	return func(s string) {
		fmt.Println(s)
		db.Store(s)
	}
}

// makeHTTPHandler build a custom handler.
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

// httpFunc is an arbitrary function type.
type httpFunc func(db DB, w http.ResponseWriter, r *http.Request) error

// handler is an arbitrary handler with the httpFunc signature.
func handler(db DB, w http.ResponseWriter, r *http.Request) error {
	return nil
}

// ExecuteFunc is an artitrary function type (may come from other party lib).
type ExecuteFunc func(string)

// Execute just executes general function with the ExecuteFunc signature.
func Execute(fn ExecuteFunc) {
	fn("foo bar baz")
}
