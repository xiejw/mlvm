package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	M = 4096
	K = M
	N = M
)

// Elapsed 12m35.339189398s
func gemm(A, B, C []float32) {
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			for k := 0; k < K; k++ {
				C[i*N+j] += A[i*K+k] * B[k*N+j]
			}
		}
	}
}

func main() {
	sizeA := M * K
	sizeB := K * N
	sizeC := M * N
	A := make([]float32, sizeA)
	B := make([]float32, sizeB)
	C := make([]float32, sizeC)

	for i := 0; i < sizeA; i++ {
		A[i] = rand.Float32()
	}

	for i := 0; i < sizeB; i++ {
		B[i] = rand.Float32()
	}

	start := time.Now()

	gemm(A, B, C)

	end := time.Now()

	fmt.Printf("Elapsed %v\n", end.Sub(start))

	for i := 0; i < sizeC; i++ {
		if i == 100 {
			break
		}
		fmt.Printf("%v ", C[i])
	}
}
