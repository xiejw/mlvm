BUILD=.build

ifdef VERBOSE
	TEST_VERBOSE=-v
endif

default: compile

compile:
	@echo "Dummy compile."

test: fmt
	go test -v github.com/xiejw/mlvm/lib/...

clean:
	rm -rf ${BUILD} ${RELEASE}

fmt:
	go mod tidy
	gofmt -w -l lib
