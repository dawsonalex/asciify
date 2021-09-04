package main

import (
	"asciify"
	"asciify/resize"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var imageFilename, outputFilename string
var outputWidth, outputHeight int

func init() {
	flag.StringVar(&imageFilename,
		"f",
		"",
		"Path to the image file to convert")
	flag.StringVar(&outputFilename,
		"o",
		"",
		"File to output art to")
	flag.IntVar(&outputWidth,
		"w",
		-1,
		"Width of the image output")
	flag.IntVar(&outputHeight,
		"g",
		-1,
		"Height of the image output")
	flag.Usage = usage
}

func main() {
	flag.Parse()
	checkArgs()
	imageBytes, err := ioutil.ReadFile(imageFilename)
	if err != nil {
		fmt.Println("error reading file: ", err)
		return
	}

	resizeOption := getResizeOptions()
	resizedImage, err := resize.Buffer(imageBytes, resizeOption)
	if err != nil {
		fmt.Printf("error resizing image: %v", err)
		os.Exit(1)
	}

	art, err := asciify.RenderFromBuffer(resizedImage, asciify.DefaultPixelMapper())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(string(art[:]))

	if len(outputFilename) > 0 {
		outputFile, err := os.Create(outputFilename)
		defer outputFile.Close()
		if err == nil {
			_, err = outputFile.Write(art)
			if err != nil {
				fmt.Printf("error writing output to file: %v", err)
			}
		} else {
			fmt.Printf("error opening file: %v", err)
		}
	}
}

func checkArgs() {
	if outputWidth != outputHeight && (outputWidth == -1 || outputHeight == -1) {
		fmt.Println("Please specify both a width and height, or neither")
		os.Exit(1)
	}
}

func getResizeOptions() resize.Option {
	if outputWidth > -1 || outputHeight > -1 {
		return resize.ToFixed(outputWidth, outputHeight)
	}

	return resize.ToTerminal()
}

func usage() {
	fmt.Fprintf(os.Stderr, `=======
asciify
=======
Usage: asciify -f <filename> [-w <width> -g <height>]
Options:
`)
	flag.PrintDefaults()
}
