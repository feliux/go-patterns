package main

import (
	"fmt"
	"sync"
	"time"
)

type DriverPg struct {
	conn string
}

var (
	instance *DriverPg
	lock     = &sync.Mutex{}
)

func Connect() *DriverPg {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
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
