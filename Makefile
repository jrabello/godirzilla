PROJECT_NAME := $(shell basename $(CURDIR))

build:
	go build -o bin/$(PROJECT_NAME)

run: build
	./bin/$(PROJECT_NAME)

clean:
	rm -f ./bin/$(PROJECT_NAME)
