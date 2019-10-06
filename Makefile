BUILD=build
FMT=gofmt -w -l

default: compile run

compile:
	go build -o ${BUILD}/main examples/main.go

run:
	./${BUILD}/main

test:
	go test -v github.com/xiejw/mlvm/mlvm/...

clean:
	rm -rf ${BUILD}

fmt:
	${FMT} mlvm && ${FMT} examples
