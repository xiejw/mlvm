/tmp/matmul: /tmp/libgemm.a
	go build -a -o /tmp/matmul matmul.go

/tmp/libgemm.a: /tmp/gemm.o
	ar -r /tmp/libgemm.a /tmp/gemm.o


/tmp/gemm.o: gemm.c gemm.h
	clang -ffast-math -march=native -O2 -c gemm.c -o /tmp/gemm.o

