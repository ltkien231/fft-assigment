package main

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	fft2 "github.com/mjibson/go-dsp/fft"
	fft1 "scientificgo.org/fft"
)

func Test_DFT(t *testing.T) {
	max := 8
	x := make([]complex128, max)
	for i := 0; i < max; i++ {
		x[i] = complex(rand.NormFloat64(), rand.NormFloat64())
		fmt.Println(x[i])
	}

	fmt.Println("\n============= My DFT ===========")
	resDFT := MyDFT(x)
	for i := 0; i < max; i++ {
		fmt.Println(resDFT[i])
	}

	fmt.Println("\n============= Lib1 FFT ===========")
	lib1FFT := fft1.Fft(x, false)
	for i := 0; i < max; i++ {
		fmt.Println(lib1FFT[i])
	}

	fmt.Println("\n============= My FFT ===========")
	resFFT := MyFFT(x)
	for i := 0; i < max; i++ {
		fmt.Println(resFFT[i])
	}
}

func Benchmark_DFT(b *testing.B) {
	powers := []int{4, 10, 15}
	for _, size := range powers {
		inputSize := int(math.Pow(2, float64(size)))
		b.Run(fmt.Sprintf("input_size_%d", inputSize), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				input := make([]complex128, inputSize)
				for i := 0; i < inputSize; i++ {
					input[i] = complex(rand.NormFloat64(), rand.NormFloat64())
				}
				MyDFT(input)
			}
		})
	}
}

func Benchmark_MyFFT(b *testing.B) {
	powers := []int{4, 10, 15}
	for _, size := range powers {
		inputSize := int(math.Pow(2, float64(size)))
		b.Run(fmt.Sprintf("input_size_%d", inputSize), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				input := make([]complex128, inputSize)
				for i := 0; i < inputSize; i++ {
					input[i] = complex(rand.NormFloat64(), rand.NormFloat64())
				}
				MyFFT(input)
			}
		})
	}
}

func Benchmark_LibFFT(b *testing.B) {
	powers := []int{4, 10, 15}
	for _, size := range powers {
		inputSize := int(math.Pow(2, float64(size)))
		b.Run(fmt.Sprintf("input_size_%d", inputSize), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				input := make([]complex128, inputSize)
				for i := 0; i < inputSize; i++ {
					input[i] = complex(rand.NormFloat64(), rand.NormFloat64())
				}
				fft1.Fft(input, false)
			}
		})
	}
}

func Benchmark_Lib2_FFT(b *testing.B) {
	powers := []int{4, 10, 15}
	for _, size := range powers {
		inputSize := int(math.Pow(2, float64(size)))
		b.Run(fmt.Sprintf("input_size_%d", inputSize), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				input := make([]complex128, inputSize)
				for i := 0; i < inputSize; i++ {
					input[i] = complex(rand.NormFloat64(), rand.NormFloat64())
				}
				fft2.FFT(input)
			}
		})
	}
}
