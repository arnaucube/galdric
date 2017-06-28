package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	readConfig("./config.json")

	c.Cyan("reading images datasets")
	tStart := time.Now()
	dataset := readDataset("./dataset")
	fmt.Print("time spend reading images: ")
	fmt.Println(time.Since(tStart))
	fmt.Println("total folders scanned: " + strconv.Itoa(len(dataset)))

	numImages := 0
	for _, v := range dataset {
		numImages = numImages + len(v)
	}
	c.Cyan("total images in dataset: " + strconv.Itoa(numImages))

	//we have the images in the dataset variable
	//now, can take images
	testFile := readImage("./test.jpg")
	r := knn(dataset, testFile)
	fmt.Println("seems to be a " + r)

}
