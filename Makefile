REPO=mlvm
LIB_DIR=go
CMD_DIR=cmd
TEST_DIR=tests
BUILD_DIR=.build
INTEGRATION_TEST=no

# folders
LIBS=github.com/xiejw/${REPO}/${LIB_DIR}/...
CMD_LIBS=github.com/xiejw/${REPO}/${CMD_DIR}/...
TEST_LIBS=github.com/xiejw/${REPO}/${TEST_DIR}/...

# cmds. convention is cmd/<binary>/main.go
CMD_CANDIDATES=$(patsubst cmd/%,%,$(wildcard cmd/*))

# verbose
TEST_FLAGS=
ifeq (1, $(VERBOSE))
TEST_FLAGS="-v"
endif

# actions
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
ifeq ($(INTEGRATION_TEST),yes)
	go fmt ${TEST_LIBS}
endif

test:
	go test $(TEST_FLAGS) ${LIBS}
ifeq ($(INTEGRATION_TEST),yes)
	go test ${TEST_LIBS}
endif

bench:
	go test -bench=. ${LIBS}

clean:
	go mod tidy
	@echo "clean '"${BUILD_DIR}"'" && rm -rf ${BUILD_DIR}

# Optionally include a local Makefile.
-include Makefile.local

C_FMT=docker run --rm -ti \
			--user `id -u ${USER}`:`id -g ${USER}` \
			-v `pwd`:/workdir xiejw/clang-format \
			/clang-format.sh

CFLAGS=-std=c99 -O2 -Wall -Werror -pedantic -Wno-c11-extensions ${CLAGS}

c:
	${CC} ${CLAGS} -Isrc -o /tmp/1 cmd/c/main.c

c_fmt:
	${C_FMT} cmd/c src

