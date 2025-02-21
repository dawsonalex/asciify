GO_MODULE_PATH = asciify

# ROOT_DIR is the path of the makefile (including trailing slash)
ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))
PROJECT_PATH := $(ROOT_DIR:/=)
BIN_NAME = asciify

# SHELL is set to a wrapper file that sources a helpers file and executes
# the command passed by Make.
#SHELL := $(PWD)/shell

.PHONY: help # Generate list of targets with descriptions
help:
	grep '^.PHONY: .* #' makefile | sed 's/\.PHONY: \(.*\) # \(.*\)/\1    \2/' | expand -t20

.PHONY: build # Build the binary (currently depends on vips)
build:
	go build -v -tags vips -o '${ROOT_DIR}${BIN_NAME}' '${GO_MODULE_PATH}/cmd'

.PHONY: run # Build and run the binary bin/imageservice
run: build
	${ROOT_DIR}/${BIN_NAME}

.PHONY: test # run all tests
test:
	go test ${GO_MODULE_PATH}/...

.PHONY: clean # remove build files
clean:
	rm -rv '${ROOT_DIR}${BIN_NAME}'

