package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	once     sync.Once
	instance *DriverPg // global variable
)

type DriverPg struct {
	conn string
}

func Connect() *DriverPg {
	once.Do(func() {
		instance = &DriverPg{conn: "DriverConnectPostgres"}
	})
	return instance
}

func main() {
	// Simulate a delayed call to Connect.
	go func() {
		time.Sleep(time.Millisecond * 600)
		fmt.Println(*Connect())
	}()
	for i := 0; i < 100; i++ {
		go func(ix int) {
			time.Sleep(time.Millisecond * 60)
			fmt.Println(ix, " = ", Connect().conn)
		}(i)
	}
	fmt.Scanln()
}
