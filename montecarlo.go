package main

import (
	"fmt"
	"math/rand"
	"math"
	"time"
)

const interval = 10000

func monte_carlo() {
	
	var circle_pt float64 = 0
	var sq_pt float64 = 0
	var dist, pi float64

	rand.Seed(time.Now().UTC().UnixNano())
	
	for i := 0; i < (interval * interval); i++ {
		// generate random x and y
		rand_x := math.Mod(rand.Float64(), interval)
		rand_y := math.Mod(rand.Float64(), interval)

		// calculate distance from origin
		dist = rand_x * rand_x + rand_y * rand_y;

		// if inside circle
		if (dist <= 1) {
			circle_pt++
		}

		// keep track of square points
		sq_pt++

		pi = float64(4 * circle_pt) / sq_pt;

	}
	fmt.Printf("Final estimate of pi: %e\n", pi)
}

func main() {
	monte_carlo()
}
