GOOS:=linux


$(shell mkdir -p build/_output/)

.PHONY: dep 
dep:
	dep ensure

.PHONY: build
build: build/_output/

build/_output/: 
	GOOS=$(GOOS) go build \
	  -a -o $@ cmd/main.go

clean:
	rm -rf build/_output/