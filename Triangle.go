// Solution to codility problem Triangle: https://codility.com/programmers/task/triangle/
package main

import "fmt"
import "sort"



func triangle(a []int) int {
	
	alen := len(a)
	if (alen < 3){
		return 0;
	}
	b := make([]int, len(a))
	copy(b, a[:])
	sort.Ints(b)
	
    for count := 0; count < (alen - 2); count++  {
        if(b[count+2] < (b[count] + b[count+1])) {
            return 1;
        }
    }
    

	return 0
}




func main() {
	fmt.Println("Codility Triangle problem")
	a := []int {10, 2, 5, 1, 8, 20}
  	
  	var result int
    result = triangle(a)
    fmt.Println("The triangule triplet exists for this array is: " , result)
}
