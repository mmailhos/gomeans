package gomeans

import (
	"math/rand"
	"strconv"
	"time"
)

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

func initClusters(k int) (clusters []Cluster) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < k; i++ {
		clusters = append(clusters, Cluster{Point{rand.Float64(), rand.Float64()}, []Point{}})
	}
	return
}

func repositionCenters(clusters []Cluster) {
	for i := 0; i < len(clusters); i++ {
		clusters[i].repositionCenter()
	}
}

/*Run runs the k-means algorithm given an array of coordinates and a specific k. Returns a slice of Clusters defined
by their Center (type Point) and a slice of Points representing points in that cluster.*/
func Run(dataset []Point, k int) []Cluster {
	pointsClusterIndex := make([]int, len(dataset))
	clusters := initClusters(k)

	for hasChanged := true; hasChanged; {
		hasChanged = false
		for i := 0; i < len(dataset); i++ {
			var minDist float64
			var updatedClusterIndex int
			for j := 0; j < len(clusters); j++ {
				tmpDist := dataset[i].Distance(clusters[j].Center)
				if minDist == 0 || tmpDist < minDist {
					minDist = tmpDist
					updatedClusterIndex = j
				}
			}
			clusters[updatedClusterIndex].Points = append(clusters[updatedClusterIndex].Points, dataset[i])
			if pointsClusterIndex[i] != updatedClusterIndex {
				pointsClusterIndex[i] = updatedClusterIndex
				hasChanged = true
			}
		}
		if hasChanged {
			repositionCenters(clusters)
		}
	}
	return clusters
}

/*RunWithDrawing runs the k-means algorithm given an array of coordinates and a specific k. Output charts in
chart folder*/
func RunWithDrawing(dataset []Point, k int) []Cluster {
	pointsClusterIndex := make([]int, len(dataset))
	clusters := initClusters(k)
	hasChanged := true

	draw(clusters, "charts/initial_centers.png")
	for p := 0; hasChanged; p++ {
		hasChanged = false
		for i := 0; i < len(dataset); i++ {
			var minDist float64
			var updatedClusterIndex int
			for j := 0; j < len(clusters); j++ {
				tmpDist := dataset[i].Distance(clusters[j].Center)
				if minDist == 0 || tmpDist < minDist {
					minDist = tmpDist
					updatedClusterIndex = j
				}
			}
			clusters[updatedClusterIndex].Points = append(clusters[updatedClusterIndex].Points, dataset[i])
			if pointsClusterIndex[i] != updatedClusterIndex {
				pointsClusterIndex[i] = updatedClusterIndex
				hasChanged = true
			}
		}
		draw(clusters, "charts/"+strconv.Itoa(p)+".png")
		if hasChanged {
			repositionCenters(clusters)
		}
	}
	return clusters
}
