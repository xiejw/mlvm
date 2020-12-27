# ------------------------------------------------------------------------------
# configurations.
# ------------------------------------------------------------------------------
SRC           = src
CMD           = cmd
BUILD_BASE    = .build
BUILD         = ${BUILD_BASE}
BUILD_RELEASE = ${BUILD_BASE}_release
UNAME         = $(shell uname)

CFLAGS        := -std=c99 -Wall -Werror -pedantic -Wno-c11-extensions ${CFLAGS}
CFLAGS        := ${CFLAGS} -I${SRC} -I../eva/src
LDFLAGS       := -lm ${LDFLAGS}

# enable POSIX
ifeq ($(UNAME), Linux)
CFLAGS := ${CFLAGS} -D_POSIX_C_SOURCE=201410L
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

FMT = docker run --rm -ti \
    --user `id -u ${USER}`:`id -g ${USER}` \
    -v `pwd`:/workdir xiejw/clang-format \
    /clang-format.sh

# ------------------------------------------------------------------------------
# libs.
# ------------------------------------------------------------------------------

ALL_LIBS =

# ------------------------------------------------------------------------------
# tests.
# ------------------------------------------------------------------------------
ADT_TEST_SUITE  = ${BUILD}/adt_vec_test.o ${BUILD}/adt_sds_test.o \
							    ${BUILD}/adt_map_test.o
ADT_TEST_DEP    = ${ADT_LIB} ${BASE_LIB}
ADT_TEST        = ${ADT_TEST_SUITE} ${ADT_TEST_DEP}

# ------------------------------------------------------------------------------
# actions.
# ------------------------------------------------------------------------------
compile: ${BUILD} ${ALL_LIBS}

${BUILD}:
	@mkdir -p ${BUILD}

${BUILD}/base_%.o: ${SRC}/base/%.c
	${EVA_CC} -o $@ -c $<

clean:
	rm -rf ${BUILD_BASE}*

fmt:
	${FMT} ${SRC}
	${FMT} ${CMD}

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

${BUILD}/mlvm: cmd/mlvm/main.c ../eva/.build_release/libeva.a
	${EVA_LD} -o $@ $^

test: compile ${BUILD}/test
	${EVA_EX} ${BUILD}/test

${BUILD}/test: cmd/test/main.c ${ADT_TEST} ${CRON_TEST} ${RNG_TEST} \
	             ${ALGORITHMS_TEST}
	${EVA_LD} -o $@ $^
