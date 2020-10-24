CFLAGS=-std=c99 -O2 -Wall -Werror -pedantic -Wno-c11-extensions

BIN=.build

${BIN}/main: ${BIN} ../cmd/c/main.c ${BIN}/opcode.o
	${CC} ${CFLAGS} -I. -o ${BIN}/main ../cmd/c/main.c ${BIN}/opcode.o

${BIN}/opcode.o: opcode.c opcode.h
	${CC} ${CFLAGS} -I. -o ${BIN}/opcode.o -c opcode.c

${BIN}:
	mkdir -p ${BIN}
