VERSION=$(shell git rev-parse --short=9 HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

GITHUB_USERNAME=nshttpd
BUILD_BASE_DIR=${GOPATH}/src/github.com/${GITHUB_USERNAME}

.PHONY : tenor-root tenor-search

all: tenor-root tenor-search

tenor-root : darwin-tenor-root linux-tenor-root arm-tenor-root
tenor-search: darwin-tenor-search linux-tenor-search arm-tenor-search

darwin-% :
	$(eval BASE=$(shell echo $@ | awk -F- '{print $$3}'))
	cd ${BASE}; \
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o tenor-${BASE}_darwin .

linux-%:
	$(eval BASE=$(shell echo $@ | awk -F- '{print $$3}'))
	cd ${BASE}; \
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o tenor-${BASE}_linux .

arm-%:
	$(eval BASE=$(shell echo $@ | awk -F- '{print $$3}'))
	cd ${BASE}; \
	CGO_ENABLED=0 GOOS=linux GOARM=7 GOARCH=arm go build -o tenor-${BASE}_arm .
