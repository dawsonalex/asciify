package main

import (
	"fmt"
	"image2ascii/asciiart"
	"io/ioutil"
)

func main() {
	fmt.Println("Hello, World")

	imageBytes, err := ioutil.ReadFile("/home/ad/go/src/image2ascii/cmd/image.png")
	if err != nil {
		fmt.Println("error reading file: ", err)
		return
	}

	art, err := asciiart.FromImageBuffer(10, 1, imageBytes)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(string(art[:]))
}
