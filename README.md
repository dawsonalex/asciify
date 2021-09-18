# ASCII art generator

Generate ASCII art from images.

## Build

``shell
go build -tags vips -o asciify cmd/main.go
``

## Run

`asciify -h` for help.

```
=======
asciify
=======
Usage: asciify -f <filename> [-w <width> -g <height>]
Options:
  -f string
        Path to the image file to convert
  -g int
        Height of the image output (default -1)
  -o string
        File to output art to
  -w int
        Width of the image output (default -1)

```