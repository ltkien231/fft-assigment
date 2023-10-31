# Fast Fourier Transform implentation

This repo implements a Fast Fourier Transform (FFT) base on the idea of recursive. Additionally, we demonstrates how FFT is applied in image compression

Structure:
- /fft_implementation:
    - dft.go: implement Discrete Fourier Transform (DFT)
    - fft.go: implement FFT
    - fft_test.go: test the rightness and performance of our implementation against the existing implementation 
- /image_compression:
    - compress.py: image compression program 
    - /images: includes original image and the compression result with three different reserve percentage
