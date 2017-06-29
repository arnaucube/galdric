package main

import (
	"fmt"
	"sort"
)

func euclideanDist(img1, img2 [][]float64) float64 {
	var dist float64
	for i := 0; i < len(img1); i++ {
		for j := 0; j < len(img1[i]); j++ {
			dist += (img1[i][j] - img2[i][j]) * (img1[i][j] - img2[i][j])
		}
	}

	return dist
}

type Neighbour struct {
	Dist  float64
	Label string
}

func isNeighbour(neighbours []Neighbour, dist float64, label string) []Neighbour {
	var temp []Neighbour

	for i := 0; i < len(neighbours); i++ {
		temp = append(temp, neighbours[i])
	}
	ntemp := Neighbour{dist, label}
	temp = append(temp, ntemp)

	//now, sort the temp array
	sort.Slice(temp, func(i, j int) bool {
		return temp[i].Dist < temp[j].Dist
	})

	for i := 0; i < len(neighbours); i++ {
		neighbours[i] = temp[i]
	}

	return neighbours
}

func knn(dataset Dataset, input [][]float64) string {
	k := 3
	var neighbours []Neighbour
	//d := euclideanDist(dataset["leopard"][0], input)
	for i := 0; i < k; i++ {
		/*neighbours[i].Dist = euclideanDist(dataset["leopard"][0], input)
		neighbours[i].Label = "leopard"*/
		neighbours = append(neighbours, Neighbour{euclideanDist(dataset["leopard"][0], input), "leopard"})
	}
	for l, v := range dataset {
		for i := 0; i < len(v); i++ {
			dNew := euclideanDist(v[i], input)
			/*if dNew < d {
				d = dNew
				label = l
			}*/
			neighbours = isNeighbour(neighbours, dNew, l)
		}
	}
	for i := 0; i < len(neighbours); i++ {
		fmt.Print(neighbours[i].Label + " - ")
		fmt.Println(neighbours[i].Dist)
	}

	return neighbours[0].Label
}
