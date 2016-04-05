package main

import "fmt"
import "math"


// Solution to codility lesson: https://codility.com/programmers/task/tape_equilibrium/

func Abs64(x int64) int64 {
    if x < 0 {
        return -x
    }
    return x
}

func tapeEquilibrium(a []int64) int64 {
	var sum int64
	var lsum int64
	var rsum int64
	var min int64  
	min = math.MaxInt64
	sum, lsum, rsum = 0, 0, 0 
	
	
	for _, v := range a {
		sum = sum + v
	}
	
	
	for _, v := range a {
		rsum = sum - lsum - v
		lsum = lsum + v
		if(min > Abs64(rsum - lsum)) {
			min = Abs64(rsum - lsum)
		}
	}
	
	return min
}




func main() {
	fmt.Println("Codility Tape Equilibrium problem")
	a := []int64 {3, 1, 2, 4, 3} 
  	
  	var result int64
    result = tapeEquilibrium(a)
    fmt.Println("The result is:" , result)
}
