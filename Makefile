# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
BINARY_NAME=api
BINARY_UNIX=$(BINARY_NAME)_unix
SRC=src/main.go

all: build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(SRC)
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
deps:
	$(GOGET) github.com/jinzhu/gorm
	$(GOGET) github.com/jinzhu/gorm/dialects/mysql

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v $(SRC)
