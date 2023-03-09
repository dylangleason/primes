IMAGE = "primes:latest"

default: build run

.PHONY: build clean run test

build:
	docker build -t $(IMAGE) .

clean:
	docker image rm $(IMAGE)

run:
	docker run --rm $(IMAGE)

test:
	docker run --rm $(IMAGE) /usr/local/go/bin/go test -v

bench:
	docker run --rm $(IMAGE) /usr/local/go/bin/go test -bench=PrimesUpTo
