/tmp/matmul: libgemm.a
	go build -a -o /tmp/matmul matmul.go

libgemm.a: gemm.o
	ar -r libgemm.a gemm.o


gemm.o: gemm.c gemm.h
	clang -ffast-math -march=native -O2 -c gemm.c

