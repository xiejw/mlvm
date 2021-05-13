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

CFLAGS          += -I${EVA_PATH}/src -g
LDFLAGS         += ${EVA_LIB}

TEX             = docker run --rm -v `pwd`:/workdir xiejw/tex pdftex

# ------------------------------------------------------------------------------
# libs.
# ------------------------------------------------------------------------------
VM_HEADER       = ${SRC}/vm.h ${SRC}/opcode.h
VM_LIB          = ${BUILD}/vm_vm.o ${BUILD}/vm_shape.o ${BUILD}/vm_tensor.o \
                  ${BUILD}/vm_primitives.o

ALL_LIBS        = ${VM_LIB}

# ------------------------------------------------------------------------------
# actions.
# ------------------------------------------------------------------------------

.DEFAULT_GOAL   = compile_cmd # mnist # regression  # vm

compile: ${BUILD} ${ALL_LIBS}

compile_cmd: compile ${BUILD}/vm ${BUILD}/regression ${BUILD}/mnist

${BUILD}/vm_%.o: ${SRC}/%.c ${VM_HEADER}
	${EVA_CC} -o $@ -c $<

# header dependencies.
${BUILD}/vm_vm.o: ${SRC}/primitives.h ${SRC}/vm_internal.h

${BUILD}/vm_tensor.o: ${SRC}/vm_internal.h

${BUILD}/vm_primitives.o: ${SRC}/primitives.h


# ------------------------------------------------------------------------------
# cmd.
# ------------------------------------------------------------------------------

.PNONY: vm

vm: compile ${BUILD}/vm
	${EVA_EX} ${BUILD}/vm

${BUILD}/vm: cmd/vm/main.c ${VM_LIB}
	${EVA_LD} -o $@ $^

regression: compile ${BUILD}/regression
	${EVA_EX} ${BUILD}/regression

${BUILD}/regression: cmd/regression/main.c ${VM_LIB}
	${EVA_LD} -o $@ $^

mnist: compile ${BUILD}/mnist
	${EVA_EX} ${BUILD}/mnist

${BUILD}/mnist: cmd/mnist/main.c ${VM_LIB}
	${EVA_LD} -o $@ $^

# ------------------------------------------------------------------------------
# docs.
# ------------------------------------------------------------------------------
#
DOC             = doc

DOCS            = ${DOC}/design.pdf ${DOC}/loss_softmax_crossentropy.pdf

.PNONY: ${DOCS}

doc: ${DOCS}

${DOC}/%.pdf: ${DOC}/%.tex
	${TEX} -output-directory `dirname "$@"` $<

