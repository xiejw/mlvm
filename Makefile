REPO=mlvm
LIB_DIR=go
CMD_DIR=cmd
BUILD_DIR=.build

LIBS=github.com/xiejw/${REPO}/${LIB_DIR}/...
CMD_LIBS=github.com/xiejw/${REPO}/${CMD_DIR}/...
CMD_CANDIDATES=$(patsubst cmd/%,%,$(wildcard cmd/*))  # convention is cmd/<binary>/main.go

compile: compile_lib compile_cmd

compile_lib:
	go build ${LIBS}

compile_cmd:
	@mkdir -p ${BUILD_DIR}
	@for cmd in ${CMD_CANDIDATES}; do \
		echo 'compile cmd/'$$cmd && \
	  go build -o ${BUILD_DIR}/$$cmd cmd/$$cmd/main.go; \
	done

fmt:
	go fmt ${LIBS}
	go fmt ${CMD_LIBS}

test:
	go test ${LIBS}

bench:
	go test -bench=. ${LIBS}

clean:
	go mod tidy
	@echo "clean '"${BUILD_DIR}"'" && rm -rf ${BUILD_DIR}

# Optionally include a local Makefile.
-include Makefile.local
