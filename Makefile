IMAGE = "primes:latest"

.PHONY: build clean run test bench

default: build run

build:
	docker build -t $(IMAGE) .

clean:
	docker image rm $(IMAGE)

run:
	docker run -it --rm $(IMAGE)

test:
	docker run --rm $(IMAGE) /usr/local/go/bin/go test -v

bench:
	docker run --rm $(IMAGE) /usr/local/go/bin/go test -bench=PrimesUpTo
