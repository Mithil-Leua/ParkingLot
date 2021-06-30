$SRC_PKGS=$(shell go list ./... | grep -v "vendor" | grep -v "load_tests")

build: compile test

compile:
	go build ../

test:
	go test ../src/$(SRC_PKGS)

run: compile
		./parkingLot
