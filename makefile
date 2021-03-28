EVA_PATH        = ../eva
EVA_LIB         = ${EVA_PATH}/.build_release/libeva.a

include ${EVA_PATH}/eva.mk

# ------------------------------------------------------------------------------
# configurations.
# ------------------------------------------------------------------------------

SRC             =  src
CMD             =  cmd
CFLAGS          += -I${SRC}

FMT_FOLDERS     =  ${SRC} ${CMD}  # required by eva.mk

CFLAGS          += -I${EVA_PATH}/src
LDFLAGS         += ${EVA_LIB}

# ------------------------------------------------------------------------------
# libs.
# ------------------------------------------------------------------------------
VM_HEADER       = ${SRC}/vm.h
VM_LIB          = ${BUILD}/vm_vm.o ${BUILD}/vm_shape.o ${BUILD}/vm_tensor.o

ALL_LIBS        = ${VM_LIB}

# ------------------------------------------------------------------------------
# actions.
# ------------------------------------------------------------------------------

.DEFAULT_GOAL   = vm

compile: ${BUILD} ${ALL_LIBS}

${BUILD}/vm_%.o: ${SRC}/%.c ${VM_HEADER}
	${EVA_CC} -o $@ -c $<

# header dependencies.
${BUILD}/vm_tensor.o: ${SRC}/vm_internal.h

# ------------------------------------------------------------------------------
# cmd.
# ------------------------------------------------------------------------------

.PNONY: vm

vm: compile ${BUILD}/vm
	${EVA_EX} ${BUILD}/vm

${BUILD}/vm: cmd/vm/main.c ${VM_LIB}
	${EVA_LD} -o $@ $^
