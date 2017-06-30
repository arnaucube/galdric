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

func getMapKey(dataset map[string]ImgDataset) string {
	for k, _ := range dataset {
		return k
	}
	return ""
}

type LabelCount struct {
	Label string
	Count int
}

func averageLabel(neighbours []Neighbour) string {
	labels := make(map[string]int)
	for _, n := range neighbours {
		labels[n.Label]++
	}
	//create array from map
	var a []LabelCount
	for k, v := range labels {
		a = append(a, LabelCount{k, v})
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i].Count > a[j].Count
	})
	fmt.Println(a)
	//send the most appeared neighbour in k
	return a[0].Label
}
func knn(dataset Dataset, input [][]float64) string {
	k := 10
	var neighbours []Neighbour
	label := getMapKey(dataset)
	for i := 0; i < k; i++ {
		/*neighbours[i].Dist = euclideanDist(dataset["leopard"][0], input)
		neighbours[i].Label = "leopard"*/
		neighbours = append(neighbours, Neighbour{euclideanDist(dataset[label][0], input), label})
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

	r := averageLabel(neighbours)
	return r
}
