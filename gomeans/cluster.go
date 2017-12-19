package gomeans

//Cluster struct composed of a center Point{X,Y float64} and of Elements
type Cluster struct {
	Center   Point
	Elements []Element
}

func (cluster *Cluster) elementsXValues() (elementsXValues []float64) {
	for i := 0; i < len(cluster.Elements); i++ {
		elementsXValues = append(elementsXValues, cluster.Elements[i].Coordinate.X)
	}
	return
}

func (cluster *Cluster) elementsYValues() (elementsYValues []float64) {
	for i := 0; i < len(cluster.Elements); i++ {
		elementsYValues = append(elementsYValues, cluster.Elements[i].Coordinate.Y)
	}
	return
}

func (cluster *Cluster) repositionCenter() {
	var clusterCount = len(cluster.Elements)
	var x, y float64

	for i := 0; i < clusterCount; i++ {
		x = x + cluster.Elements[i].Coordinate.X
		y = y + cluster.Elements[i].Coordinate.Y
	}
	cluster.Elements = []Element{}
	cluster.Center = Point{x / float64(clusterCount), y / float64(clusterCount)}
}
