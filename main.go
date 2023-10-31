package main

import (
	"fmt"
	"math"
)

func main() {
	omega := 2 * math.Pi * 64
	omegaSincos := math.Cos(omega) - math.Sin(omega)
	fmt.Println("omegaSincos", math.Pow(omegaSincos, 3))

}
