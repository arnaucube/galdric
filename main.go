package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

//dataset := make(Dataset)
var dataset Dataset
var datasetED Dataset

func main() {
	readConfig("./config.json")

	c.Cyan("reading images datasets")
	tStart := time.Now()
	dataset, datasetED = readDataset("./dataset")
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

	c.Green("server running")
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":"+config.ServerPort, router))
}
