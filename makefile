EVA_PATH        = ../eva
EVA_LIB         = ${EVA_PATH}/.build_release/libeva.a

include ${EVA_PATH}/eva.mk

# ------------------------------------------------------------------------------
# Configurations.
# ------------------------------------------------------------------------------

SRC             =  src
CMD             =  cmd
CFLAGS          += -I${SRC}

FMT_FOLDERS     =  ${SRC} ${CMD}  # required by eva.mk

CFLAGS          += -I${EVA_PATH}/src -g
LDFLAGS         += ${EVA_LIB}

TEX             = docker run --rm -v `pwd`:/workdir xiejw/tex pdftex

CMDS            = $(patsubst ${CMD}/%,%,$(wildcard ${CMD}/*))
CMD_TARGETS     = $(patsubst ${CMD}/%/main.c,${BUILD}/%,$(wildcard ${CMD}/*/main.c))

# ------------------------------------------------------------------------------
# Libs.
# ------------------------------------------------------------------------------
VM_HEADER       = ${SRC}/vm.h ${SRC}/opcode.h
VM_LIB          = ${BUILD}/vm_vm.o ${BUILD}/vm_shape.o ${BUILD}/vm_tensor.o \
                  ${BUILD}/vm_primitives.o

ALL_LIBS        = ${VM_LIB}

ifdef BLIS
CFLAGS  += -DBLIS=1 -I../blis/include/${BLIS}/ -Wno-unused-function
LDFLAGS += ../blis/lib/${BLIS}/libblis.a -pthread
endif

# ------------------------------------------------------------------------------
# Header Deps.
# ------------------------------------------------------------------------------
${BUILD}/vm_vm.o: ${SRC}/primitives.h ${SRC}/vm_internal.h

${BUILD}/vm_tensor.o: ${SRC}/vm_internal.h

${BUILD}/vm_primitives.o: ${SRC}/primitives.h

# ------------------------------------------------------------------------------
# Actions.
# ------------------------------------------------------------------------------

.DEFAULT_GOAL   = compile

compile: ${BUILD} ${ALL_LIBS}

${BUILD}/vm_%.o: ${SRC}/%.c ${VM_HEADER}
	${EVA_CC} -o $@ -c $<

# ------------------------------------------------------------------------------
# Cmd.
# ------------------------------------------------------------------------------

compile: ${CMD_TARGETS}

$(foreach cmd,$(CMDS),$(eval $(call objs,$(cmd),$(BUILD),$(VM_LIB))))

# ------------------------------------------------------------------------------
# Docs.
# ------------------------------------------------------------------------------

DOC             = doc
DOCS            = ${DOC}/design.pdf ${DOC}/loss_softmax_crossentropy.pdf

doc: ${DOCS}

${DOC}/%.pdf: ${DOC}/%.tex
	${TEX} -output-directory `dirname "$@"` $<
