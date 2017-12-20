package gomeans

import (
	"bytes"
	"github.com/wcharczuk/go-chart"
	"io/ioutil"
)

func draw(clusters []Cluster, outfile string) {
	var series []chart.Series

	for i := 0; i < len(clusters); i++ {
		series = append(series, chart.ContinuousSeries{
			Style: chart.Style{
				Show:        true,
				StrokeWidth: chart.Disabled,
				DotWidth:    5},
			XValues: clusters[i].pointsXValues(),
			YValues: clusters[i].pointsYValues(),
		})
	}
	series = append(series, chart.ContinuousSeries{
		Style: chart.Style{
			Show:        true,
			StrokeWidth: chart.Disabled,
			DotWidth:    5,
		},
		XValues: centersX(clusters),
		YValues: centersY(clusters),
	})
	graph := chart.Chart{
		Series: series,
		XAxis: chart.XAxis{
			Style: chart.Style{
				Show: true,
			},
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true,
			},
		},
	}

	buffer := bytes.NewBuffer([]byte{})
	graph.Render(chart.PNG, buffer)
	ioutil.WriteFile(outfile, buffer.Bytes(), 0644)
}
