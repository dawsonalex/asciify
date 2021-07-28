package main

import (
	"fmt"
	"image2ascii/asciiart"
	"io/ioutil"
)

func main() {
	fmt.Println("Hello, World")

	imageBytes, err := ioutil.ReadFile("/home/ad/GolandProjects/image2ascii/cmd/image.jpeg")
	if err != nil {
		fmt.Println("error reading file: ", err)
		return
	}

	asciiart.FromImageBuffer(10, 1, imageBytes)
}
