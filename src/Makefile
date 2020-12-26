CFLAGS :=-std=c99 -Wall -Werror -pedantic -Wno-c11-extensions ${CFLAGS}
CFLAGS :=${CFLAGS} -I.

BUILD_BASE=.build
BUILD=${BUILD_BASE}

ifeq (1, $(RELEASE))
  CFLAGS :=${CFLAGS} -DNDEBUG -O2
  BUILD :=${BUILD}_release
endif

ifeq (1, $(ASAN))
  CFLAGS :=${CFLAGS} -g -fsanitize=address
  BUILD :=${BUILD}_asan
endif


OPCODE=${BUILD}/opcode.o
OPCODE_TEST=${BUILD}/opcode_test.o
VEC=${BUILD}/vec.o
VEC_TEST=${BUILD}/vec_test.o
PROGRAM=${BUILD}/program.o
MLVM_H=mlvm.h
SDS=${BUILD}/sds.o
SDS_TEST=${BUILD}/sds_test.o

MAIN=main.c
MAIN_LIBS=${OPCODE} ${VEC} ${SDS} ${PROGRAM}

TEST=test.c
TEST_H=test.h
TEST_SUITES=${VEC_TEST} ${OPCODE_TEST} ${SDS_TEST}
TEST_LIBS=${VEC} ${OPCODE} ${SDS} ${PROGRAM}

default: ${BUILD}/main
	@echo "compiled as ${BUILD}/main"

${OPCODE}: opcode.c opcode.h ${MLVM_H}
	${CC} ${CFLAGS} -o $@ -c $<

${PROGRAM}: program.c program.h ${MLVM_H}
	${CC} ${CFLAGS} -o $@ -c $<

${VEC}: vec.c vec.h ${MLVM_H}
	${CC} ${CFLAGS} -o $@ -c $<

${SDS}: sds.c sds.h ${MLVM_H}
	${CC} ${CFLAGS} -o $@ -c $<

${BUILD}:
	mkdir -p ${BUILD}

${BUILD}/main: ${BUILD} ${MAIN} ${MAIN_LIBS}
	${CC} ${CFLAGS} -o ${BUILD}/main ${MAIN} ${MAIN_LIBS}

test: ${BUILD}/test
	${BUILD}/test

${BUILD}/test: ${BUILD} ${TEST} ${TEST_H} ${TEST_LIBS} ${TEST_SUITES}
	${CC} ${CFLAGS} -o ${BUILD}/test ${TEST} ${TEST_LIBS} ${TEST_SUITES}

${VEC_TEST}: vec_test.c ${VEC}
	${CC} ${CFLAGS} -o $@ -c $<

${OPCODE_TEST}: opcode_test.c ${OPCODE}
	${CC} ${CFLAGS} -o $@ -c $<

${SDS_TEST}: sds_test.c ${SDS}
	${CC} ${CFLAGS} -o $@ -c $<

clean:
	rm -rf ${BUILD_BASE}*

run: ${BUILD}/main
	${BUILD}/main
