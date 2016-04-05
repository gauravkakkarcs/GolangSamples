// Solution to codility problem MaxDoubleSliceSum: https://codility.com/programmers/task/max_double_slice_sum/
package main

import "fmt"

func Max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func maxDoubleSliceSum(a []int) int {
	alen := len(a)
	if (alen == 3){
		return 0;
	}
	L := make([]int, alen)
	R := make([]int, alen)

    for count := 1; count < alen-1; count++  {
        L[count] = Max(0, a[count] + L[count - 1])
    }
    
    for count := alen-2; count > 0; count-- {
        R[count] = Max(0, a[count] + R[count + 1])
    }

    max := 0
    for count := 1; count < alen-1; count++ {
        if(max < (L[count-1] + R[count+1])) {
        	max = (L[count-1] + R[count+1]);
        }
    }
    
    return max
}




func main() {
	fmt.Println("Codility MaxDoubleSliceSum problem")
	a := []int {3, 2, 6, -1, 4, 5, -1, 2}
  	
  	var result int
    result = maxDoubleSliceSum(a)
    fmt.Println("The maximal sum of any double slice is:" , result)
}


