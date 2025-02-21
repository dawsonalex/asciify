package main

import (
	"asciify"
	"asciify/resize"
	"flag"
	"fmt"
	"os"
	"runtime/debug"
	"strings"
)

var imageFilename, outputFilename string
var outputWidth, outputHeight int
var invertLightness bool

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
	flag.BoolVar(&invertLightness,
		"i",
		false,
		"Invert the lightness of output")
	flag.Usage = usage
}

func main() {
	flag.Parse()
	checkArgs()
	imageBytes, err := os.ReadFile(imageFilename)
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

	charset := asciify.CharsetDarkToLight
	if invertLightness {
		charset = asciify.CharsetLightToDark
	}
	art, err := asciify.RenderFromBuffer(resizedImage, asciify.CharSetMapper(charset))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(string(art[:]))

	if len(outputFilename) > 0 {
		outputFile, err := os.Create(outputFilename)
		defer func(outputFile *os.File) {
			err := outputFile.Close()
			if err != nil {
				fmt.Printf("error closing %s: %v", outputFilename, err)
			}
		}(outputFile)
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
		_, _ = fmt.Println("Please specify both a width and height, or neither")
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
	buildStr := "[NO BUILD INFO]"
	if bi, ok := debug.ReadBuildInfo(); ok {
		if bi.Main.Version[len(bi.Main.Version)-5:] == "dirty" {
			buildStr = bi.Main.Version
		} else {
			buildStr = strings.Split(bi.Main.Version, "-")[0]
		}
	}

	_, _ = fmt.Fprintf(os.Stderr, `=======
asciify %s
=======
Usage: asciify -f <filename> [-w <width> -g <height>]
Options:
`, buildStr)
	flag.PrintDefaults()
}
