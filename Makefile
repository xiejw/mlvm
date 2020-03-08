DEBUG=.debug
RELEASE=./.release

FMT=docker run --rm -ti \
      --user `id -u ${USER}`:`id -g ${USER}` \
      -v `pwd`:/source xiejw/clang-format \
      /clang-format.sh

.PHONY: default compile compile_only run test fmt clean

default: run

compile:
	mkdir -p ${DEBUG} && cd ${DEBUG} && CLICOLOR_FORCE=1 cmake -GNinja .. && ninja

compile_only:
	@echo "-> Compiling..." && cd ${DEBUG} && ninja

run: compile_only
	@echo "-> Running..." && ${DEBUG}/compile

test: compile
	SKIP_LONG_TEST=1 ${DEBUG}/test

fmt:
	@echo "-> Formatting..." && ${FMT} cmd

clean:
	@echo "-> Cleaning..." && rm -rf ${DEBUG} ${RELEASE}
