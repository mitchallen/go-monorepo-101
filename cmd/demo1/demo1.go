package main

import (
	"fmt"

	"github.com/mitchallen/coin"
	"github.com/mitchallen/go-monorepo-101/pkg/alpha"
	"github.com/mitchallen/go-monorepo-101/pkg/beta"
)

func DemoFlip() {

	// Demo Flip()

	m := map[bool]int{
		true:  0,
		false: 0,
	}

	for i := 0; i < 100; i++ {
		m[coin.Flip()]++
	}

	fmt.Print("[Flip]: ")
	fmt.Println(m)

	// Use alpha too:

	total := alpha.Add(m[true], m[false])

	fmt.Printf("Total: %d \n", total)

	// use beta
	b1 := beta.Multiply(m[true], m[false])
	fmt.Printf("b1: %d \n", b1)
}

func main() {
	DemoFlip()
}
