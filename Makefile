# -----------------------------------------------------------------------------
# configurations.
# -----------------------------------------------------------------------------

C_FMT=docker run --rm -ti \
			--user `id -u ${USER}`:`id -g ${USER}` \
			-v `pwd`:/workdir xiejw/clang-format \
			/clang-format.sh
TEST_FLAGS=

ifeq (1, $(VERBOSE))
  TEST_FLAGS="-v"
endif

# -----------------------------------------------------------------------------
# actions.
# -----------------------------------------------------------------------------

compile: c

c:
	make -j -C src

release: clean
	RELEASE=1 make -j -C src

test:
	make -j -C src test

asan:
	ASAN=1 make -C src test

fmt:
	@echo "fmt src" && ${C_FMT} src

clean:
	make -C src clean

run:
	make -j -C src run
