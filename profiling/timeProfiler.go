package main

import (
	"fmt"
	"log"
	"math/big"
	"time"
)

// Duration calculates the duration time.
func Duration(invocation time.Time, name string) {
	elapsed := time.Since(invocation)
	log.Printf("%s lasted %s", name, elapsed)
}

// BigIntFactorial calculates the factorial of a big int number.
func BigIntFactorial(x *big.Int) *big.Int {
	// Arguments to a defer statement is immediately evaluated and stored.
	// The deferred function receives the pre-evaluated values when its invoked.
	defer Duration(time.Now(), "IntFactorial")
	y := big.NewInt(1)
	for one := big.NewInt(1); x.Sign() > 0; x.Sub(x, one) {
		y.Mul(y, x)
	}
	return x.Set(y)
}

func main() {
	x := big.NewInt(10)
	result := BigIntFactorial(x)
	fmt.Printf("factorial is %d", result)
}
