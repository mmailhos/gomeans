package gomeans

import (
	"math/rand"
	"strconv"
	"time"
)

//Element struct represents a single point. It has coordinates (Point{X,Y float64}) and is part of a cluster
type Element struct {
	Coordinate Point
	ClusterID  int
}

func centersX(clusters []Cluster) (centersX []float64) {
	for i := 0; i < len(clusters); i++ {
		centersX = append(centersX, clusters[i].Center.X)
	}
	return
}

func centersY(clusters []Cluster) (centersY []float64) {
	for i := 0; i < len(clusters); i++ {
		centersY = append(centersY, clusters[i].Center.Y)
	}
	return
}

func elementsFromPoints(initialDataset []Point) (dataset []Element) {
	for i := 0; i < len(initialDataset); i++ {
		dataset = append(dataset, Element{Coordinate: initialDataset[i], ClusterID: 0})
	}
	return
}

func initClusters(k int) (clusters []Cluster) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < k; i++ {
		clusters = append(clusters, Cluster{Point{rand.Float64(), rand.Float64()}, []Element{}})
	}
	return
}

func iterate(dataset []Element, clusters []Cluster) ([]Cluster, bool) {
	hasChanged := false
	for i := 0; i < len(dataset); i++ {
		var minDist float64
		var updatedClusterIndex int
		for j := 0; j < len(clusters); j++ {
			tmpDist := dataset[i].Coordinate.Distance(clusters[j].Center)
			if minDist == 0 || tmpDist < minDist {
				minDist = tmpDist
				updatedClusterIndex = j
			}
		}
		clusters[updatedClusterIndex].Elements = append(clusters[updatedClusterIndex].Elements, dataset[i])
		if dataset[i].ClusterID != updatedClusterIndex {
			dataset[i].ClusterID = updatedClusterIndex
			hasChanged = true
		}
	}
	return clusters, hasChanged
}


/*RunWithDrawing runs the k-means algorithm given an array of coordinates and a specific k. Output is some charts in
chart folder*/
func RunWithDrawing(initialDataset []Point, k int) []Cluster {
	dataset := elementsFromPoints(initialDataset)
	clusters := initClusters(k)
	hasChanged := true

	draw(clusters, "charts/initial_centers.png")
	for p := 0; hasChanged; p++ {
		hasChanged = false
		for i := 0; i < len(dataset); i++ {
			var minDist float64
			var updatedClusterIndex int
			for j := 0; j < len(clusters); j++ {
				tmpDist := dataset[i].Coordinate.Distance(clusters[j].Center)
				if minDist == 0 || tmpDist < minDist {
					minDist = tmpDist
					updatedClusterIndex = j
				}
			}
			clusters[updatedClusterIndex].Elements = append(clusters[updatedClusterIndex].Elements, dataset[i])
			if dataset[i].ClusterID != updatedClusterIndex {
				dataset[i].ClusterID = updatedClusterIndex
				hasChanged = true
			}
		}
		draw(clusters, "charts/"+strconv.Itoa(p)+".png")
		for i := 0; i < len(clusters); i++ {
			clusters[i].repositionCenter()
		}
	}
	return clusters
}
