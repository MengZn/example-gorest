GOOS:=linux
REPO := quay.io
OWNER := mengzn
APP:=example-gorest

DOCKERFILE:=./build

VERSION:=1.0

.PHONY: dep 
dep:
	dep ensure

.PHONY: build
build: build/_output/bin/rest

build/_output/bin/rest:
	$(shell mkdir -p build/_output/bin/)
	GOOS=$(GOOS) go build \
	  -a -o $@

.PHONY: build_image
build_image:
	docker build $(DOCKERFILE) --tag $(REPO)/$(OWNER)/$(APP):$(VERSION)

.PHONY: push_image
push_image:
	docker push $(REPO)/$(OWNER)/$(APP):$(VERSION)

clean:
	rm -rf build/_output/