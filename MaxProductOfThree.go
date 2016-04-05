// Solution to codility problem MaxProductOfThree: https://codility.com/programmers/task/max_product_of_three/
package main

import "fmt"
import "sort"



func maxProductOfThree(a []int) int {
	
	alen := len(a)
	if (alen < 3){
		return 0;
	}
	
	sort.Ints(a)
	max := a[alen-1] * a[alen-2] * a[alen-3]
	
    if(max < a[0]*a[1]*a[alen-1]) {
        max = a[0]*a[1]*a[alen-1]
    }else if(max < a[0]*a[alen-2]*a[alen-1]) {
    	max = a[0]*a[alen-2]*a[alen-1]
    }
        
	
	return max
}




func main() {
	fmt.Println("Codility MaxProductOfThree problem")
	a := []int {-3, 1, 2, -2, 5, 6}
  	
  	var result int
    result = maxProductOfThree(a)
    fmt.Println("The maximal product of any triplet is:" , result)
}
