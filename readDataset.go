package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

//each image is [][]float64, is a array of pixels
type ImgDataset [][][]float64

type Dataset map[string]ImgDataset

func byteArrayToFloat64Array(b []byte) []float64 {
	var f []float64
	for i := 0; i < len(b); i++ {
		val, _ := strconv.ParseFloat(string(b[i]), 64)
		f = append(f, val)
	}
	return f
}

func readImage(path string) [][]float64 {
	//open image file
	dat, err := ioutil.ReadFile(path)
	check(err)

	pathSplited := strings.Split(path, ".")
	imageExtension := pathSplited[len(pathSplited)-1]
	imageRaw, err := dataToImage(dat, imageExtension)
	check(err)

	//resize the image to standard size
	image := Resize(imageRaw)

	//convert the image to histogram(RGBA)
	histogram := imageToHistogram(image)
	return histogram
}
func readDataset(path string) map[string]ImgDataset {
	dataset := make(Dataset)

	folders, _ := ioutil.ReadDir(path)
	for _, folder := range folders {
		fmt.Println(folder.Name())

		var imgDataset ImgDataset

		folderFiles, _ := ioutil.ReadDir(path + "/" + folder.Name())
		for _, file := range folderFiles {
			image := readImage(path + "/" + folder.Name() + "/" + file.Name())

			imgDataset = append(imgDataset, image)

		}

		//add the foldername to the Dataset map
		dataset[folder.Name()] = imgDataset
	}

	return dataset
}
