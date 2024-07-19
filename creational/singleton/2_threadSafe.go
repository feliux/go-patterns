package main

import (
	"fmt"
	"sync"
	"time"
)

// DriverPg is the driver for PSQL.
type DriverPg struct {
	conn string
}

var (
	instance *DriverPg
	lock     = &sync.Mutex{}
)

// Connect just connects to the database.
func Connect() *DriverPg {
	// Lock here is so aggresive cause the following if evaluation
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		instance = &DriverPg{conn: "DriverConnectPostgres"}
	}
	return instance
}

func main() {
	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(time.Millisecond * 600)
			fmt.Println(*Connect(), " - ", i)
		}
	}()
	go func() {
		fmt.Println(*Connect())
	}()
	fmt.Scanln()
}
