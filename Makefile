# ------------------------------------------------------------------------------
# configurations.
# ------------------------------------------------------------------------------
SRC           = src
CMD           = cmd
BUILD_BASE    = .build
BUILD         = ${BUILD_BASE}
BUILD_RELEASE = ${BUILD_BASE}_release
UNAME         = $(shell uname)

EVA_LIB       = ../eva/.build_release/libeva.a

CFLAGS        := -std=c99 -Wall -Werror -pedantic -Wno-c11-extensions ${CFLAGS}
CFLAGS        := ${CFLAGS} -I${SRC} -I../eva/src
LDFLAGS       := -lm ${LDFLAGS} ${EVA_LIB}
MK            := make

# enable POSIX
ifeq ($(UNAME), Linux)
CFLAGS := ${CFLAGS} -D_POSIX_C_SOURCE=201410L
endif

ifeq ($(UNAME), FreeBSD)
MK := gmake
endif

# enable asan by `make ASAN=1`
ifdef ASAN
	CFLAGS := ${CFLAGS} -g -fsanitize=address -D_ASAN=1
	BUILD  := ${BUILD}_asan
endif

# enable release by `make RELEASE=1`
ifdef RELEASE
  CFLAGS := ${CFLAGS} -DNDEBUG -O2
  BUILD  := ${BUILD}_release

compile: check_release_folder
endif

FMT = clang-format -i --style=file

# compact print with some colors.
EVA_CC    = ${QUIET_CC}${CC} ${CFLAGS}
EVA_LD    = ${QUIET_LD}${CC} ${LDFLAGS} ${CFLAGS}
EVA_AR    = ${QUIET_AR}ar -cr
EVA_EX    = ${QUIET_EX}

CCCOLOR   = "\033[34m"
LINKCOLOR = "\033[34;1m"
SRCCOLOR  = "\033[33m"
BINCOLOR  = "\033[36;1m"
ENDCOLOR  = "\033[0m"

# enable verbose cmd by `make V=1`
ifndef V
QUIET_CC  = @printf '    %b %b\n' $(CCCOLOR)CC$(ENDCOLOR) \
				  $(SRCCOLOR)`basename $@`$(ENDCOLOR) 1>&2;
QUIET_LD  = @printf '    %b %b\n' $(LINKCOLOR)LD$(ENDCOLOR) \
				  $(BINCOLOR)`basename $@`$(ENDCOLOR) 1>&2;
QUIET_AR  = @printf '    %b %b\n' $(LINKCOLOR)AR$(ENDCOLOR) \
				  $(BINCOLOR)`basename $@`$(ENDCOLOR) 1>&2;
QUIET_EX  = @printf '    %b %b\n' $(LINKCOLOR)EX$(ENDCOLOR) \
				  $(BINCOLOR)$@$(ENDCOLOR) 1>&2;
endif


# ------------------------------------------------------------------------------
# libs.
# ------------------------------------------------------------------------------
VM_LIB = ${BUILD}/vm_opcode.o

ALL_LIBS = ${VM_LIB}

# ------------------------------------------------------------------------------
# tests.
# ------------------------------------------------------------------------------
VM_TEST_SUITE  = ${BUILD}/vm_opcode_test.o
VM_TEST_DEP    = ${VM_LIB}
VM_TEST        = ${VM_TEST_SUITE} ${VM_TEST_DEP}

ALL_TESTS      = ${VM_TEST}

# ------------------------------------------------------------------------------
# actions.
# ------------------------------------------------------------------------------
compile: ${BUILD} ${ALL_LIBS} ${EVA_LIB}

${EVA_LIB}:
	${MK} -C ../eva RELEASE=1 libeva

${BUILD}:
	@mkdir -p ${BUILD}

${BUILD}/vm_%.o: ${SRC}/vm/%.c
	${EVA_CC} -o $@ -c $<

clean:
	rm -rf ${BUILD_BASE}*

fmt:
	find ${SRC} ${CMD} -iname *.h -o -iname *.c | xargs ${FMT}

check_release_folder:
ifneq (${BUILD}, ${BUILD_RELEASE})
	@echo "release mode cannot mix with other modes, e.g., asan."
	@exit 1
endif

# ------------------------------------------------------------------------------
# cmds.
# ------------------------------------------------------------------------------

mlvm: compile ${BUILD}/mlvm
	${EVA_EX} ${BUILD}/mlvm

${BUILD}/mlvm: cmd/mlvm/main.c
	${EVA_LD} -o $@ $^

test: compile ${BUILD}/test
	${EVA_EX} ${BUILD}/test

${BUILD}/test: cmd/test/main.c ${ALL_TESTS}
	${EVA_LD} -o $@ $^
