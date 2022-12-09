
# Use eg:
#   $ make run ARGS="4 1"

build: bin/aoc2021

clean:
	rm bin/*

run: build
	./bin/aoc2021 $(ARGS)

bin/aoc2021:  $(shell find . -type f -name "*.go")  $(shell find input -type f -name "*.txt")
	go build -v -o bin/aoc2021
