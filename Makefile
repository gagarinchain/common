# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=gnetwork
BINARY_UNIX=$(BINARY_NAME)_unix
all: test build
protos:
	cd protobuff/protos && PATH=$(PATH):$(GOPATH)/bin protoc --go_out=$(PKGMAP):.. *.proto