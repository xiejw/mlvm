DEBUG=.debug
RELEASE=.release

FMT=docker run --rm -ti \
      --user `id -u ${USER}`:`id -g ${USER}` \
      -v `pwd`:/workdir xiejw/clang-format \
      /clang-format.sh

.PHONY: default compile compile_only run test doc fmt clean

default: prng

compile:
	@echo "-> Bootstraping..."
	mkdir -p ${DEBUG} && cd ${DEBUG} && cmake -DMLVM_ADDRESS_SANITIZER=OFF -GNinja .. && ninja -v

compile_only:
	@echo "-> Compiling..." && cd ${DEBUG} && ninja -v

prng: compile_only
	@echo "-> Running..." && ${DEBUG}/prng

test: compile_only
	@echo "-> Testing..." && ${DEBUG}/test

test_address_sanitizer: clean
	mkdir -p ${DEBUG} && cd ${DEBUG} && cmake -DMLVM_ADDRESS_SANITIZER=ON -GNinja .. && ninja -v && ./test

doc:
	make -C doc

fmt:
	@echo "-> Formatting..." && ${FMT} cmd mlvm

clean:
	@echo "-> Cleaning..." && rm -rf ${DEBUG} ${RELEASE} && make -C doc clean
