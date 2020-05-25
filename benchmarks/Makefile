FOLDER=/tmp/matmul_test

FMT=docker run --rm -ti \
      --user `id -u ${USER}`:`id -g ${USER}` \
      -v `pwd`:/workdir xiejw/clang-format \
      /clang-format.sh

${FOLDER}/matmul: ${FOLDER}/libgemm.a matmul.go
	go build -a -o ${FOLDER}/matmul matmul.go

${FOLDER}/libgemm.a: ${FOLDER}/gemm.o
	ar -r ${FOLDER}/libgemm.a ${FOLDER}/gemm.o


${FOLDER}/gemm.o: gemm.c gemm.h ${FOLDER}
	clang -ffast-math -march=native -O2 -I${HOME}/Workspace/blis/include/haswell/ -c gemm.c -o ${FOLDER}/gemm.o

${FOLDER}:
	mkdir ${FOLDER}

clean:
	rm -rf ${FOLDER}/

fmt:
	${FMT} .
