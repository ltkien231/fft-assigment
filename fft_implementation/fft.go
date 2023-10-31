package main

func MyFFT(x []complex128) []complex128 {
	N := len(x)
	if N <= 1 {
		return x
	}

	y := make([]complex128, N)

	// base case
	if N == 2 {
		// [ 1 1 ]  * [ x0 ]
		// [ 1 w ]    [ x1 ]
		y[0] = x[0] + x[1]
		y[1] = x[0] + omega(2, 1, 1)*x[1]
	}

	// recursively process into 2 sub vector
	xEven := make([]complex128, N/2)
	xOdd := make([]complex128, N/2)
	for k := 0; k < N/2; k++ {
		xEven[k] = x[2*k]
		xOdd[k] = x[2*k+1]
	}

	resultEven := MyFFT(xEven)
	resultOdd := MyFFT(xOdd)

	//Reconstruct the result: X(k) = Even(k) + W(N/SIZE*k%N) * Odd(k)
	for k := 0; k < N/2; k++ {
		y[k] = resultEven[k] + omega(N, 1, k)*resultOdd[k]
		y[k+N/2] = resultEven[k] - omega(N, 1, k)*resultOdd[k]
	}

	return y
}
