package main

import (
	"fmt"
	"math/rand"
	"math"
	"time"
)

func main() {
	const interval = 10000
	var circle_pt float64 = 0
	var sq_pt float64 = 0
	var dist, pi float64

	rand.Seed(time.Now().UTC().UnixNano())
	
	for i := 0; i < (interval * interval); i++ {
		rand_x := math.Mod(rand.Float64(), interval)
		rand_y := math.Mod(rand.Float64(), interval)

		dist = rand_x * rand_x + rand_y * rand_y;

		if (dist <= 1) {
			circle_pt++
		}

		sq_pt++

		pi = float64(4 * circle_pt) / sq_pt;

	}
	fmt.Println("Final estimate of pi: %e", pi)
}
