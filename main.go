package main

import (
	"./gomeans"
	"flag"
	"math/rand"
	"os"
	"time"
)

func main() {
	var dataset []gomeans.Point
	rand.Seed(time.Now().UnixNano())

	k := flag.Int("k", 0, "number of clusters")
	size := flag.Int("n", 0, "number of elements")
	flag.Parse()

	if *k == 0 || *size == 0 {
		flag.Usage()
		os.Exit(1)
	}

	// Generate some random points
	for i := 0; i < *size; i++ {
		dataset = append(dataset, gomeans.Point{rand.Float64(), rand.Float64()})
	}

	//Runs and outputs charts for 4 clusters
	gomeans.RunWithDrawing(dataset, *k)
}
