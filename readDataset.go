package main

import (
	"fmt"
	"image"
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

func readImage(path string) image.Image {
	//open image file
	dat, err := ioutil.ReadFile(path)
	check(err)

	pathSplited := strings.Split(path, ".")
	imageExtension := pathSplited[len(pathSplited)-1]
	imageRaw, err := dataToImage(dat, imageExtension)
	check(err)

	//resize the image to standard size
	image := Resize(imageRaw)

	return image
	/*
		//convert the image to histogram(RGBA)
		histogram := imageToHistogram(image)
		return histogram
	*/
}
func readDataset(path string) []Dataset {
	var resultDatasets []Dataset
	dataset := make(Dataset)
	datasetED := make(Dataset)
	/*
		datasetG := make(Dataset)
	*/

	folders, _ := ioutil.ReadDir(path)
	for _, folder := range folders {
		fmt.Println(folder.Name())

		var imgDataset ImgDataset
		var imgDatasetED ImgDataset
		/*
			var imgDatasetG ImgDataset
		*/

		folderFiles, _ := ioutil.ReadDir(path + "/" + folder.Name())
		for _, file := range folderFiles {
			//get the image as original
			image := readImage(path + "/" + folder.Name() + "/" + file.Name())
			histogram := imageToHistogram(image)
			imgDataset = append(imgDataset, histogram)

			//get the image with EdgeDetection filter
			imageED := EdgeDetection(image)
			histogramED := imageToHistogram(imageED)
			imgDatasetED = append(imgDatasetED, histogramED)

			/*
				//get the image with Grayscale filter
				imageG := Grayscale(image)
				histogramG := imageToHistogram(imageG)
				imgDatasetG = append(imgDatasetG, histogramG)
			*/
		}

		//add the foldername to the Dataset map
		dataset[folder.Name()] = imgDataset
		datasetED[folder.Name()] = imgDatasetED
		/*
			datasetG[folder.Name()] = imgDatasetG
		*/
	}

	resultDatasets = append(resultDatasets, dataset)
	resultDatasets = append(resultDatasets, datasetED)
	/*
		resultDatasets = append(resultDatasets, datasetG)
	*/
	return resultDatasets
}
