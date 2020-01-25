BUILD=.build

ifdef VERBOSE
	TEST_VERBOSE=-v
endif

default: compile

init:
	mkdir -p ${BUILD}

compile: init
	go build -o ${BUILD}/hello cmd/main.go

run:
	${BUILD}/hello

test: fmt
	go test ${TEST_VERBOSE} github.com/xiejw/mlvm/lib/...

clean:
	rm -rf ${BUILD} ${RELEASE}

fmt:
	go mod tidy
	gofmt -w -l lib cmd
