package main

func euclideanDist(img1, img2 [][]float64) float64 {
	var dist float64

	for i := 0; i < len(img1); i++ {
		for j := 0; j < len(img1[i]); j++ {
			dist += (img1[i][j] - img2[i][j]) * (img1[i][j] - img2[i][j])
		}
	}

	return dist
}

func knn(dataset Dataset, input [][]float64) string {
	d := euclideanDist(dataset["Leopards"][0], input)
	label := "lamp"
	for k, v := range dataset {
		//fmt.Println(k)
		for i := 0; i < len(v); i++ {
			//fmt.Println(i)
			dNew := euclideanDist(v[i], input)
			if dNew < d {
				d = dNew
				label = k
			}
		}
	}
	return label
}
