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
	mkdir -p data
	go test -v -json  -count=1 -bench=BenchmarkAll* -run=^a ./... > data/benchdata


python-tests:
	python -m unittest discover  -v -s visual/
