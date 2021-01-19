include ../eva/eva.mk

# ------------------------------------------------------------------------------
# configurations.
# ------------------------------------------------------------------------------
SRC           = src
CMD           = cmd
FMT_FOLDERS   = ${SRC} ${CMD}

EVA_LIB       = ../eva/.build_release/libeva.a

CFLAGS        += -I../eva/src
LDFLAGS       += ${EVA_LIB}

# ------------------------------------------------------------------------------
# libs.
# ------------------------------------------------------------------------------
VM_LIB = ${BUILD}/vm_opcode.o ${BUILD}/vm_object.o ${BUILD}/vm_stack.o \
	 ${BUILD}/vm_vm.o

ALL_LIBS = ${VM_LIB}

# ------------------------------------------------------------------------------
# tests.
# ------------------------------------------------------------------------------
VM_TEST_SUITE  = ${BUILD}/vm_opcode_test.o ${BUILD}/vm_object_test.o
VM_TEST_DEP    = ${VM_LIB}
VM_TEST        = ${VM_TEST_SUITE} ${VM_TEST_DEP}

ALL_TESTS      = ${VM_TEST}

# ------------------------------------------------------------------------------
# actions.
# ------------------------------------------------------------------------------

.DEFAULT_GOAL = compile

compile: ${BUILD} ${ALL_LIBS} ${EVA_LIB}

${EVA_LIB}:
	${MK} -C ../eva RELEASE=1 libeva

${BUILD}/vm_%.o: ${SRC}/vm/%.c
	${EVA_CC} -o $@ -c $<

# ------------------------------------------------------------------------------
# cmds.
# ------------------------------------------------------------------------------

m: mlvm

mlvm: compile ${BUILD}/mlvm
	${EVA_EX} ${BUILD}/mlvm

${BUILD}/mlvm: cmd/mlvm/main.c ${VM_LIB}
	${EVA_LD} -o $@ $^

test: compile ${BUILD}/test
	${EVA_EX} ${BUILD}/test

${BUILD}/test: cmd/test/main.c ${ALL_TESTS}
	${EVA_LD} -o $@ $^
