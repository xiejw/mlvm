DEBUG=./.debug
RELEASE=./.release

compile:
	mkdir -p ${DEBUG} && cd ${DEBUG} && CLICOLOR_FORCE=1 cmake .. && make -j

compile_only:
	@cd ${DEBUG} && make -j --no-print-directory

run: compile_only
	@${DEBUG}/compile

test: compile
	SKIP_LONG_TEST=1 ${DEBUG}/test

fmt:
	docker run --rm -ti \
      --user `id -u ${USER}`:`id -g ${USER}` \
      -v `pwd`:/source xiejw/clang-format \
      /clang-format.sh .

clean:
	rm -rf ${DEBUG} ${RELEASE}
