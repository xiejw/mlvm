EVA_PATH        = ../eva
EVA_LIB         = ${EVA_PATH}/.build_release/libeva.a

include ${EVA_PATH}/eva.mk

# ------------------------------------------------------------------------------
# Configurations.
# ------------------------------------------------------------------------------

SRC             =  src
INCLUDE         =  include
CMD             =  cmd
FMT_FOLDERS     =  ${SRC} ${CMD} ${INCLUDE}  # required by eva.mk

VM_SRC          =  ${SRC}/vm

CFLAGS          += -I${VM_SRC} -I${INCLUDE} -I${EVA_PATH}/src
CFLAGS          += -DVM_INTERNAL=1

ifndef RELEASE
CFLAGS          += -g
endif

LDFLAGS         += ${EVA_LIB}

TEX             = docker run --rm -v `pwd`:/workdir xiejw/tex pdftex

# ------------------------------------------------------------------------------
# Libs.
# ------------------------------------------------------------------------------
VM_HEADER       = ${INCLUDE}/vm.h ${VM_SRC}/op.h
VM_LIB          = ${BUILD}/vm_vm.o ${BUILD}/vm_shape.o ${BUILD}/vm_tensor.o \
                  ${BUILD}/vm_primitives.o

ALL_LIBS        = ${VM_LIB}

# ------------------------------------------------------------------------------
# Header Deps.
# ------------------------------------------------------------------------------
${BUILD}/vm_vm.o: ${VM_SRC}/primitives.h ${VM_SRC}/vm_internal.h

${BUILD}/vm_tensor.o: ${VM_SRC}/vm_internal.h

${BUILD}/vm_primitives.o: ${VM_SRC}/primitives.h

# ------------------------------------------------------------------------------
# Actions.
# ------------------------------------------------------------------------------

.DEFAULT_GOAL   = compile

compile: ${BUILD} ${ALL_LIBS}

${BUILD}/vm_%.o: ${VM_SRC}/%.c ${VM_HEADER}
	${EVA_CC} -o $@ -c $<

libmlvm: compile ${BUILD}/libmlvm.a

${BUILD}/libmlvm.a: ${VM_LIB}
	${EVA_AR} $@ $^

# ------------------------------------------------------------------------------
# Cmd.
# ------------------------------------------------------------------------------

# Put `test` out from CMDS, as it needs special testing library in next section.
CMD_CANDIDATES  = $(patsubst ${CMD}/%,%,$(wildcard ${CMD}/*))
CMDS            = $(filter-out test,${CMD_CANDIDATES})
CMD_TARGETS     = $(patsubst ${CMD}/%/main.c,${BUILD}/%,$(wildcard ${CMD}/*/main.c))

compile: ${CMD_TARGETS}

$(foreach cmd,$(CMDS),$(eval $(call objs,$(cmd),$(BUILD),$(VM_LIB))))

# ------------------------------------------------------------------------------
# Tests.
# ------------------------------------------------------------------------------
TEST_LIBS       = ${BUILD}/shape_test.o ${BUILD}/tensor_test.o \
		  ${BUILD}/vm_test.o ${BUILD}/op_test.o

${BUILD}/%_test.o: ${VM_SRC}/%_test.c ${VM_HEADER}
	${EVA_CC} -o $@ -c $<

$(eval $(call objs,test,$(BUILD),$(VM_LIB) $(TEST_LIBS)))


# ------------------------------------------------------------------------------
# Docs.
# ------------------------------------------------------------------------------

DOC             = doc
DOCS            = ${DOC}/design.pdf ${DOC}/loss_softmax_crossentropy.pdf

doc: ${DOCS}

${DOC}/%.pdf: ${DOC}/%.tex
	${TEX} -output-directory `dirname "$@"` $<
