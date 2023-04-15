.DEFAULT_GOAL := build


build:
	go build contest/solution.go

run:
	go run contest/solution.go

test:
	go test -v ./...

benchmark:
	go test -v -bench=BenchmarkAll* -run=^a.  ./...

graphs:
	mkdir -p data
	go test -v -json  -count=1 -bench=BenchmarkAll* -run=^a ./... > data/benchdata
	python -m visual.src.run


python-tests:
	python -m unittest discover  -v -s visual/

python-setup:
	mkdir venv
	python3.11 -m venv venv

python-depends:
	pip install -r requirements.txt
