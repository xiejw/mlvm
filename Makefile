REPO=mlvm
LIBS_DIR=vm compiler
CMD_DIR=cmd
BUILD_DIR=.build

# ------------------------------------------------------------------------------
# mappings.
# ------------------------------------------------------------------------------
# folders, where LIBs can be multiple.
LIBS=$(patsubst %,github.com/xiejw/${REPO}/%/...,${LIBS_DIR})
CMD_LIBS=github.com/xiejw/${REPO}/${CMD_DIR}/...

# cmds. convention is cmd/<binary>/main.go
CMD_CANDIDATES=$(patsubst cmd/%,%,$(wildcard cmd/*))

# verbose for testing.
TEST_FLAGS=
ifeq (1, $(VERBOSE))
TEST_FLAGS="-v"
endif

# ------------------------------------------------------------------------------
# actions.
# ------------------------------------------------------------------------------
compile: compile_libs compile_cmd

compile_libs:
	go build ${LIBS}

compile_cmd:
	@mkdir -p ${BUILD_DIR}
	@for cmd in ${CMD_CANDIDATES}; do \
		echo 'compile cmd/'$$cmd && \
	  go build -o ${BUILD_DIR}/$$cmd cmd/$$cmd/main.go; \
	done

fmt:
	go fmt ${LIBS} ${CMD_LIBS}

test:
	go test $(TEST_FLAGS) ${LIBS}

.PHONY: tags
tags:
	ctags -R ${LIBS_DIR}

clean:
	@echo "clean 'go.mod'" && go mod tidy
	@echo "clean '"${BUILD_DIR}"'" && rm -rf ${BUILD_DIR}

