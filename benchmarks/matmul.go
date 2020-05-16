package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// #cgo LDFLAGS: -L/tmp/ -lgemm
// #include <gemm.h>
import "C"

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

// Elapsed 2m8.324053419s
func gemmLoopOrderNaive(A, B, C []float32) {
	for i := 0; i < M; i++ {
		for k := 0; k < K; k++ {
			for j := 0; j < N; j++ {
				C[i*N+j] += A[i*K+k] * B[k*N+j]
			}
		}
	}
}

// Elapsed 1m31.383356293s
func gemmLoopOrder(A, B, C []float32) {
	for i := 0; i < M; i++ {
		for k := 0; k < K; k++ {
			a := A[i*K+k]
			indexC := i * N
			indexB := k * N

			for j := 0; j < N; j++ {
				C[indexC] += a * B[indexB]
				indexC++
				indexB++
			}
		}
	}
}

// Elapsed 17.097657839s
func gemmLoopOrderWithParallelism(A, B, C []float32, numThreads int) {
	wg := new(sync.WaitGroup)

	var delta int
	delta = M / numThreads
	if M%numThreads != 0 {
		delta += 1
	}

	for p := 0; p < numThreads; p++ {

		wg.Add(1)

		go func(start int) {

			startIndex := start * delta
			endIndex := startIndex + delta
			if endIndex >= M {
				endIndex = M - 1
			}

			for i := startIndex; i < endIndex; i++ {
				for k := 0; k < K; k++ {
					a := A[i*K+k]
					indexC := i * N
					indexB := k * N

					for j := 0; j < N; j++ {
						C[indexC] += a * B[indexB]
						indexC++
						indexB++
					}
				}
			}
			wg.Done()

		}(p)

	}
	wg.Wait()
}

func gemmCallCWithParallelism(A, B, D []float32, numThreads int) {
	wg := new(sync.WaitGroup)

	var delta int
	delta = M / numThreads
	if M%numThreads != 0 {
		delta += 1
	}

	for p := 0; p < numThreads; p++ {

		wg.Add(1)

		go func(start int) {

			startIndex := start * delta
			endIndex := startIndex + delta
			if endIndex >= M {
				endIndex = M - 1
			}

			result, err := C.gemm1((*C.float)(&D[0]), (*C.float)(&A[0]), (*C.float)(&B[0]), C.int(startIndex),
			C.int(endIndex))

			if err != nil {
				panic(fmt.Sprintf("hello, %v result: %v", result,  err))
			}

			wg.Done()

		}(p)

	}
	wg.Wait()
}

// Optimization Missed
// - restricted keyword in C
// - Loop unroll?
// - SIMD
// - clang flags, like avx fast-math
func main() {
	sizeA := M * K
	sizeB := K * N
	sizeC := M * N
	A := make([]float32, sizeA)
	B := make([]float32, sizeB)
	D := make([]float32, sizeC)

	for i := 0; i < sizeA; i++ {
		A[i] = rand.Float32()
	}

	for i := 0; i < sizeB; i++ {
		B[i] = rand.Float32()
	}

	numCPU := runtime.NumCPU()

	fmt.Printf("Num CPU: %v\n", numCPU)

	start := time.Now()

	// gemm(A, B, D)
	// gemmLoopOrderNaive(A, B, D)
	// gemmLoopOrder(A, B, D)
	// Disable hyper-thread https://software.intel.com/content/www/us/en/develop/articles/setting-thread-affinity-on-smt-or-ht-enabled-systems.html
	// gemmLoopOrderWithParallelism(A, B, D, numCPU/2)
	gemmCallCWithParallelism(A, B, D, numCPU/2)

	end := time.Now()
	fmt.Printf("Elapsed %v\n", end.Sub(start))

	for i := 0; i < sizeC; i++ {
		if i == 100 {
			break
		}
		fmt.Printf("%v ", D[i])
	}

	fmt.Println("")
}
