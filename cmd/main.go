package main

import (
	"asciify"
	"flag"
	"fmt"
	"github.com/qeesung/image2ascii/terminal"
	"io/ioutil"
	"os"
)

var imageFilename, outputFilename string
var fillTerminal bool
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
	flag.BoolVar(&fillTerminal,
		"t",
		true,
		"Resize the output to fill the terminal")
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

	if fillTerminal {
		t := terminal.NewTerminalAccessor()
		if width, height, err := t.ScreenSize(); err == nil {
			outputWidth = width
			outputHeight = height
		} else {
			fmt.Printf("error checking terminal size: %v\n", err)
		}
	}
	art, err := asciify.FromImageBuffer(outputWidth, outputHeight, imageBytes)
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
	if outputWidth > -1 && outputHeight > -1 {
		fillTerminal = false
	}

	if outputWidth != outputHeight && (outputWidth == -1 || outputHeight == -1) {
		fmt.Println("Please specify both a width and height, or neither")
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `=======
asciify
=======
Usage: asciify [-t] -f <filename> [-w <width> -g <height>]
Options:
`)
	flag.PrintDefaults()
}
