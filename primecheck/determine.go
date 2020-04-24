package primecheck

import (
	"fmt"
	"math"
)

//IsPrime determinesk  if the input is prime by using value mod i 
func IsPrime(value int) bool {
    for i := 2; i <= int(math.Floor(float64(value) / 2)); i++ {
        if value%i == 0 {
            return false
        }
    }
	return value > 1
}
// Primeck takes int64 and returns a b
func Primeck(number int) bool {  
	isPrime := true  
	if number == 0 || number == 1 {  
	 return false  
	} 
	 for i := 2; i <= number/2; i++ {  
	  if number%i == 0 {  
	   fmt.Printf(" %d is not a  prime number\n", number)  
	   isPrime = false  
	   break  
	  }  
	 }  
	 if isPrime == true {  
	  return isPrime 
	 } 
	 return isPrime  
   }  