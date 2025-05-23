package main

import( 
	"fmt"
	"math"
)

func isPrime(num int) bool {
    if num < 2 {
        return false
    }
    sqrtNum := int(math.Sqrt(float64(num)))
    for i := 2; i <= sqrtNum; i++ {
        if num%i == 0 {
            return false
        }
    }
    return true
}

func main() {
    number := 1931
    result := isPrime(number)
    if result {
        fmt.Printf("%d is prime\n", number)
    } else {
        fmt.Printf("%d is not prime\n", number)
    }
}
