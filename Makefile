DEBUG=.debug
RELEASE=.release

ifdef VERBOSE
	TEST_VERBOSE=-v
endif

default: compile run

compile:
	mkdir -p ${DEBUG} && cd ${DEBUG} && CLICOLOR_FORCE=1 cmake .. && make -j

run:
	${DEBUG}/example

test:
	echo "Hello"

clean:
	rm -rf ${BUILD}

fmt:
	docker run --rm -ti \
    --user `id -u ${USER}`:`id -g ${USER}` \
    -v `pwd`:/source xiejw/clang-format \
    /clang-format.sh .
