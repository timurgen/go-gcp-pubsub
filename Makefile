GOCMD=go
GOBUILD=$(GOCMD) build 
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BINARY_BASE_NAME=pubsubservice

APP_NAME = ohuenno/$(BINARY_BASE_NAME)

ifeq ($(OS),Windows_NT)
	BINARY_NAME=$(BINARY_BASE_NAME).exe
else
	BINARY_NAME=$(BINARY_BASE_NAME)
endif
    
all: deps test build
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v -ldflags="-s -w"
	docker build -t $(APP_NAME) --build-arg binaryname=$(BINARY_NAME) .
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
deps:
	$(GOGET) github.com/gorilla/mux
	$(GOGET) cloud.google.com/go/pubsub

