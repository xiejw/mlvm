DEBUG=./.debug
RELEASE=./.release

FMT=docker run --rm -ti \
      --user `id -u ${USER}`:`id -g ${USER}` \
      -v `pwd`:/source xiejw/clang-format \
      /clang-format.sh

compile:
	mkdir -p ${DEBUG} && cd ${DEBUG} && CLICOLOR_FORCE=1 cmake .. && make -j

compile_only:
	@cd ${DEBUG} && make -j --no-print-directory

run:
	@${DEBUG}/compile

test: compile
	SKIP_LONG_TEST=1 ${DEBUG}/test

fmt:
	${FMT} cmd

clean:
	rm -rf ${DEBUG} ${RELEASE}
