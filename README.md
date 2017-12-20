# Gomeans

An implementation of the k-means clustering algorithm in Golang in the cartesian coordinate plane.

### Definition
k-means clustering aims to partition _n_ observations into _k_ clusters in which each observation belongs to the cluster with the nearest mean, serving as a prototype of the cluster.

### Output
![gif](https://github.com/MathieuMailhos/gomeans/blob/master/charts/readme.gif)

Here, k=4, n=800. Each point's coordinates have been randomly generated, explaining this uniform distribution.


### Installation
You will need to have [go-chart](https://github.com/wcharczuk/go-chart) in your `$GOPATH` to create charts.
```
$ go get -u github.com/wcharczuk/go-chart
```

An example of use of the library is in `main.go` that you can simply execute as follow:
```
$ go run main.go -k 3 -n 800
```

### Documentation
There are two simple structures: a `Point`, that can be displayed using their X and Y. And a `Cluster`, defined by a center and a slice of `Point`.

```go
type Point struct {
  X float64
  Y float64
}

type Cluster struct {
  Center Point
  Points []Point
}
```

Then, there are basically two main exposed functrions.
`Run` is a simple function that returns a slice of `Cluster` given a dataset of `Point` with `k int`, the number of wanted clusters.
And `RunWithDrawing` that does the same thing except that it also generates graphs of each step of the algorithm in the `charts/` folder, as PNG.

```go
func Run(dataset []Point, k int) []Cluster
func RunWithDrawing(dataset []Point, k int) []Cluster
```

### Algorithm
It is relatively efficient: O(2tkn), where n is # objects, k is # clusters, and t  is # iterations. Normally, k, t, d << n.
Check out the references for more details about the algorithm.

### References
* [Data Clustering Algorithms](https://sites.google.com/site/dataclusteringalgorithms/k-means-clustering-algorithm)
* [K-means clustering](https://en.wikipedia.org/wiki/K-means_clustering)

### Others
Contributions, issues and questions welcomed.
