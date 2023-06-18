PROJECT_NAME := $(shell basename $(CURDIR))

build:
	go build -o bin/gdz

run: build
	./bin/gdz

clean:
	rm -f ./bin/gdz
