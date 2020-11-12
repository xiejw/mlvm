REPO=mlvm
VM_LIB_DIR=vm
COMPILER_LIB_DIR=compiler
CMD_DIR=cmd
BUILD_DIR=.build

# folders
VM_LIBS=github.com/xiejw/${REPO}/${VM_LIB_DIR}/...
COMPILER_LIBS=github.com/xiejw/${REPO}/${COMPILER_LIB_DIR}/...
CMD_LIBS=github.com/xiejw/${REPO}/${CMD_DIR}/...

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
	go build ${VM_LIBS} ${COMPILER_LIBS}

compile_cmd:
	@mkdir -p ${BUILD_DIR}
	@for cmd in ${CMD_CANDIDATES}; do \
		echo 'compile cmd/'$$cmd && \
	  go build -o ${BUILD_DIR}/$$cmd cmd/$$cmd/main.go; \
	done

fmt:
	go fmt ${VM_LIBS} ${COMPILER_LIBS} ${CMD_LIBS}

test:
	go test $(TEST_FLAGS) ${VM_LIBS} ${COMPILER_LIBS}

clean:
	go mod tidy
	@echo "clean '"${BUILD_DIR}"'" && rm -rf ${BUILD_DIR}

# Optionally include a local Makefile.
-include Makefile.local

