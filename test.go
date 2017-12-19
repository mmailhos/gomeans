package main

import (
	"./gomeans"
	"math/rand"
	"time"
)

func main() {
	var size = 800
	var dataset []gomeans.Point
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < size; i++ {
		dataset = append(dataset, gomeans.Point{rand.Float64(), rand.Float64()})
	}

	gomeans.RunWithDrawing(dataset, 2)
}
