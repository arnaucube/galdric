package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

//each image is [][]float64, is a array of pixels
type ImgDataset [][][]float64

type Dataset map[string]ImgDataset

func byteArrayToFloat64Array(b []byte) []float64 {
	var f []float64
	for i := 0; i < len(b); i++ {
		val, _ := strconv.ParseFloat(string(b[i]), 64)
		/*fmt.Print(string(b[i]) + "-")
		fmt.Println(val)*/
		f = append(f, val)
	}
	return f
}

func readImage(path string) [][]float64 {
	//open image file
	/*reader, err := os.Open(path)
	check(err)
	defer reader.Close()*/

	dat, err := ioutil.ReadFile(path)
	check(err)

	imageRaw, err := dataToImage(dat, path)
	check(err)

	//resize the image to standard size
	image := Resize(imageRaw)

	//convert the image to histogram(RGBA)
	histogram := imageToHistogram(image)
	//convert image to bytes
	/*imgBytes, err := imageToData(image, path)
	check(err)*/

	//imgFloat := byteArrayToFloat64Array(imgBytes)
	return histogram
}
func readDataset(path string) map[string]ImgDataset {
	//dataset := make(map[string]ImgDataset)
	dataset := make(Dataset)

	folders, _ := ioutil.ReadDir(path)
	for _, folder := range folders {
		fmt.Println(folder.Name())

		var imgDataset ImgDataset

		folderFiles, _ := ioutil.ReadDir(path + "/" + folder.Name())
		for _, file := range folderFiles {
			image := readImage(path + "/" + folder.Name() + "/" + file.Name())

			imgDataset = append(imgDataset, image)

			/*fmt.Println(folder.Name())
			fmt.Println(file.Name())*/
		}

		//add the foldername to the Dataset map
		dataset[folder.Name()] = imgDataset
	}

	return dataset
}
