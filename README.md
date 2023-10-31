# Fast Fourier Transform implentation

This repo implements a Fast Fourier Transform (FFT) base on the idea of recursive. Additionally, we demonstrates how FFT is applied in image compression

## /fft_implementation

Structure

- dft.go: implement Discrete Fourier Transform (DFT)
- fft.go: implement FFT
- fft_test.go: test the correctness and performance of our implementation against the existing implementation 

Run test 

```go
// test if our fft is correct
go test -run Test_FFT


// benchmark dft  
go test -bench Benchmark_DFT -run=^$ .
// benchmark our fft 
go test -bench Benchmark_MyFFT -run=^$ .
// benchmark fft from library scientificgo.org/fft
go test -bench Benchmark_LibFFT -run=^$ .
// benchmark fft from library github.com/mjibson/go-dsp/fft
go test -bench Benchmark_Lib2_FFT -run=^$ .
```

## /image_compression

Struture

- compress.py: image compression program 
- /images: includes original image and the compression result with three different reserve percentage

Run compression 

```python
python3 compression.py
```