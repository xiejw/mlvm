DEBUG=.debug
RELEASE=.release

FMT=docker run --rm -ti \
      --user `id -u ${USER}`:`id -g ${USER}` \
      -v `pwd`:/workdir xiejw/clang-format \
      /clang-format.sh

.PHONY: default compile compile_only run test fmt clean

default: prng

compile:
	@echo "-> Bootstraping..."
	mkdir -p ${DEBUG} && cd ${DEBUG} && cmake -GNinja .. && ninja

compile_only:
	@echo "-> Compiling..." && cd ${DEBUG} && ninja -v

prng: compile_only
	@echo "-> Running..." && ${DEBUG}/prng

# test: compile_only
# 	SKIP_LONG_TEST=1 ${DEBUG}/test

fmt:
	@echo "-> Formatting..." && ${FMT} cmd mlvm

clean:
	@echo "-> Cleaning..." && rm -rf ${DEBUG} ${RELEASE}
