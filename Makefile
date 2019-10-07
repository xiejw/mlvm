BUILD=build
FMT=gofmt -w -l

ifdef VERBOSE
	TEST_VERBOSE=-v
endif

default: compile run

compile:
	go build -o ${BUILD}/main examples/main.go

run:
	./${BUILD}/main

test:
	go test ${TEST_VERBOSE} github.com/xiejw/mlvm/mlvm/...

clean:
	rm -rf ${BUILD}

fmt:
	${FMT} mlvm && ${FMT} examples
