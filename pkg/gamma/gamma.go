package gamma

import (
	"github.com/mitchallen/coin"
)

func CoinCount(limit int) map[bool]int {

	// Demo Flip()

	m := map[bool]int{
		true:  0,
		false: 0,
	}

	for i := 0; i < limit; i++ {
		m[coin.Flip()]++
	}

	// fmt.Print("[Flip]: ")
	// fmt.Println(m)

	return m
}
