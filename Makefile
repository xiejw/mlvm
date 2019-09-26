default: compile run

compile:
	mkdir -p build && cd build && CLICOLOR_FORCE=1 cmake .. && make -j

clean:
	rm -rf build

fmt:
	docker run --rm -ti --user `id -u ${USER}`:`id -g ${USER}` -v `pwd`:/source xiejw/clang-format /clang-format.sh mlvm

run:
	./build/test_app

test: compile
	./build/tests
