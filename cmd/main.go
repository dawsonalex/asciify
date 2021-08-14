package main

import (
	"asciify"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	imagePath := os.Args[1]
	fmt.Printf("image path: %v\n", imagePath)
	imageBytes, err := ioutil.ReadFile(imagePath)
	if err != nil {
		fmt.Println("error reading file: ", err)
		return
	}

	art, err := asciify.FromImageBuffer(10, 1, imageBytes)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(string(art[:]))
}
