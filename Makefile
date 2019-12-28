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

release:
		mkdir -p ${RELEASE} && \
			cd ${RELEASE} && \
			CLICOLOR_FORCE=1 cmake -DCMAKE_BUILD_TYPE=RELEASE .. && \
			make -j

test: compile
	${DEBUG}/test

clean:
	rm -rf ${DEBUG} ${RELEASE}

fmt:
	docker run --rm -ti \
    --user `id -u ${USER}`:`id -g ${USER}` \
    -v `pwd`:/source xiejw/clang-format \
    /clang-format.sh .
