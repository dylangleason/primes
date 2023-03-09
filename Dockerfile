FROM golang:1.20

WORKDIR /go/src/github.com/dylangleason/primes
COPY . .

RUN go install -v ./cmd/...

CMD ["primes"]