package main

import (
	"./gomeans"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var size = 800 //Number of points
	var dataset []gomeans.Point
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < size; i++ {
		dataset = append(dataset, gomeans.Point{rand.Float64(), rand.Float64()})
	}

        //Runs and outputs charts for 4 clusters
	gomeans.RunWithDrawing(dataset, 4)

        //Runs and prints 3 clusters
	fmt.Println(gomeans.Run(dataset, 3))
}
