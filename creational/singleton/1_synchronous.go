package main

import (
	"fmt"
	"time"
)

// DriverPg is the driver for PSQL.
type DriverPg struct {
	conn string
}

var instance *DriverPg

// Connect just connects to the database.
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
