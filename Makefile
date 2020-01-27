BUILD=.build
PACKAGE=github.com/xiejw/mlvm

ifdef VERBOSE
	TEST_VERBOSE=-v
endif

default: compile

compile: init
	go build -o ${BUILD}/hello cmd/main.go

run:
	${BUILD}/hello

# {{{2 Maintainence
init:
	mkdir -p ${BUILD}

clean:
	rm -rf ${BUILD}

fmt:
	go mod tidy
	go fmt ${PACKAGE}/...

test: fmt
	go test ${TEST_VERBOSE} ${PACKAGE}/...
