.DEFAULT_GOAL := build


build:
	go build contest/solution.go

run:
	go run contest/solution.go

test:
	go test -v ./...

benchmark:
	go test -v -bench=.  ./...

graphs:
	mkdir data
	go test -v -json  -count=1 -bench=BenchmarkAll*  ./... > data/benchdata
