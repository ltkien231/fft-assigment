package main

import "math"

/*
[

j=1		Omg(0,0)    .... 	Omg(0,n-1)
j=2		....				....
j=n-1	Omg(0,n-1)  .... 	Omg(n-1,n-1)

]
*/
func MyDFT(x []complex128) []complex128 {
	n := len(x)

	res := make([]complex128, n)
	for j := 0; j < n; j++ {
		for k := 0; k < n; k++ {
			res[j] += x[k] * omega(n, j, k)
		}
	}

	return res
}

func omega(n, j, k int) complex128 {
	sin, cos := math.Sincos(-2 * math.Pi * float64(k*j) / float64(n))
	return complex(cos, sin)
}
