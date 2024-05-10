package main

import (
	"fmt"
	"time"
)

type DriverPg struct {
	conn string
}

var instance *DriverPg

func Connect() *DriverPg {
	if instance == nil {
		// Not thread safe when using goroutines
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
