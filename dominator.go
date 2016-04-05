// Solution to codility problem Dominator: https://codility.com/programmers/task/dominator/
package main

import "fmt"




func dominator(a []int) int {
	
	alen := len(a)
	if (alen == 0){
		return -1;
	}
	if (alen == 1){
		return 0;
	}
	count, index, occur, lastidx := 0, 0, 0, 0
	var maj int
	
	maj = a[index] 
	count = 1
	
    for index = 1; index < alen; index++  {
        if(maj != a[index] && count > 1) {
            count--
        } else if(maj != a[index] && count == 1) {
            maj = a[index]
            count = 1
        } else if(maj == a[index]) {
            count++
        }
    }
    
    for index = 0; index < alen; index++  {
        if(maj == a[index]) {
            occur++
            lastidx = index;           
        }
    }
    
    compare := (alen + 1)>>1
    if(alen & 0x01 == 1) {
    	compare = alen >> 1
    }
    
    if(occur > compare) {
        return lastidx
    }
    
    return -1
	
}




func main() {
	fmt.Println("Codility Dominator problem")
	a := []int {3, 4, 3, 2, 3, -1, 3, 3}
  	
  	var result int
    result = dominator(a)
    fmt.Println("The  index of dominator is:" , result)
}


