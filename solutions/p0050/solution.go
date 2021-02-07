package solution

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}
	res := 1.0
	for n > 0 {
		if n&1 == 1 {
			res *= x
		}
		x *= x
		n >>= 1
	}
	return res
}

// func myPow(x float64, n int) float64 {
//     if ( n == 0 ){
//         return(1)
//     }
//     if ( n < 0 ) {
//         return(1.0/myPow(x,-n))
//     }
//     if( n % 2 != 0 ) {
//         return(x * myPow(x,n-1))
//     }
//     return(myPow(x*x,n/2))
// }
