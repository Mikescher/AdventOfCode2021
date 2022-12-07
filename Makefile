
build: bin/aoc2021

clean:
	rm bin/*

run: build
	./bin/aoc2021 $(ARGS)

bin/aoc2021:  $(shell find . -type f -name "*.go")
	go build -v -o bin/aoc2021
