package main

import (
	"alpha"
	"fmt"

	"github.com/mitchallen/coin"
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
}

func main() {
	DemoFlip()
}
