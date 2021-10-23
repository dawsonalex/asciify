BINARY ?= $(shell basename "$(PWD)")# binary name
CMD := $(wildcard cmd/*.go)
temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))

# Clean the build directory (before committing code, for example)
.PHONY: clean
clean:
	rm -rv bin

PLATFORMS := linux/amd64

release: $(PLATFORMS)

$(PLATFORMS):
	GOOS=$(os) GOARCH=$(arch) go build -tags vips -o 'bin/$(BINARY)-$(os)-$(arch)' $(CMD)

.PHONY: release $(PLATFORMS)

