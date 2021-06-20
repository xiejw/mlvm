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

CFLAGS          += -I${SRC} -I${INCLUDE} -I${EVA_PATH}/src -g
LDFLAGS         += ${EVA_LIB}

TEX             = docker run --rm -v `pwd`:/workdir xiejw/tex pdftex

# ------------------------------------------------------------------------------
# Libs.
# ------------------------------------------------------------------------------
VM_HEADER       = ${INCLUDE}/vm.h ${SRC}/op.h
VM_LIB          = ${BUILD}/vm_vm.o ${BUILD}/vm_shape.o ${BUILD}/vm_tensor.o \
                  ${BUILD}/vm_primitives.o

ALL_LIBS        = ${VM_LIB}

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

libmlvm: compile ${BUILD}/libmlvm.a

${BUILD}/libmlvm.a: ${ALL_LIBS}
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

${BUILD}/%_test.o: ${SRC}/%_test.c ${VM_HEADER}
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
