package main

import (
  "fmt"
  "math"
	"math/cmplx"
)


// Abdul Rehman

//import libraries as not importing so copied
func fft(a []complex128, n int) []complex128 {
	x := make([]complex128, n)
	copy(x, a)

	j := 0
	for i := 0; i < n; i++ {
		if i < j {
			x[i], x[j] = x[j], x[i]
		}
		m := n / 2
		for {
			if j < m {
				break
			}
			j = j - m
			m = m / 2
			if m < 2 {
				break
			}
		}
		j = j + m
	}
	kmax := 1
	for {
		if kmax >= n {
			return x
		}
		istep := kmax * 2
		for k := 0; k < kmax; k++ {
			theta := complex(0.0, -1.0*math.Pi*float64(k)/float64(kmax))
			for i := k; i < n; i += istep {
				j := i + kmax
				temp := x[j] * cmplx.Exp(theta)
				x[j] = x[i] - temp
				x[i] = x[i] + temp
			}
		}
		kmax = istep
	}
}
func FFT(x []complex128, n int) []complex128 {
	y := fft(x, n)
	for i := 0; i < n; i++ {
		y[i] = y[i] / complex(float64(n), 0.0)
	}
	return y
}


func fourier_series(){
x0 := []float64{
		5,
		32,
		38,
		-33,
		-19,
		-10,
		1,
		-8,
		-20,
		10,
		-1,
		4,
		11,
		-1,
		-7,
		-2,
	}
  n := len(x0)
	x := make([]complex128, n)
	for k := 0; k < n; k++ {
		x[k] = complex(x0[k], 0.0)
	}

	y := FFT(x, n)
	

	fmt.Println(" K   DATA  FOURIER TRANSFORM  ")
	for k := 0; k < n ; k++ {
		fmt.Printf("%2d %6.1f  %8.3f%8.3f   \n",
			k, x0[k], real(y[k]), imag(y[k]))
	}
  }

//end Abdul rehman
var choice int
//--------- my function --mustafain
func slope_formula_2(a, b float64)float64{
  c := math.Pow(a+b, 2)
  return c

}

func slope(){
  var n1 float64
  var n2 float64
  fmt.Println("Enter Your First number: ")
  fmt.Scanln(&n1)
  fmt.Println("Enter Second number: ")
  fmt.Scanln(&n2)
  fmt.Printf("Result : %.2f", slope_formula_2(n1,n2))
  }
//---------

  
func main() {

fmt.Println("Enter your choice in number: ")
fmt.Scanln(&choice)


switch choice {
case 1:
    fmt.Println("Monday")
case 2:
    fmt.Println("Tuesday")
case 3:
    slope()
case 4:
    fmt.Println("Thursday")
case 5:
    fmt.Println("Friday")
case 6:
    fmt.Println("Saturday")
case 7:
    fmt.Println("Sunday")
default:
    fmt.Println("Invalid")
}
}
