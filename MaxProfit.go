// Solution to codility problem MaxProfit: https://codility.com/programmers/task/max_profit/
package main

import "fmt"

func Max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func Min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func maxProfit(a []int) int {
	alen := len(a)
	if (alen == 0 || alen == 1){
		return 0;
	}
	
	// Kadane Algorithm
	maxsofar := 0
    maxendinghere := 0
    minprice := a[0]
    
    for i := 1; i < alen; i++ {
    	
        maxendinghere = Max(0, a[i] - minprice);
        minprice = Min(a[i], minprice);
        maxsofar = Max(maxendinghere, maxsofar);
    }
    
    return maxsofar;
}




func main() {
	fmt.Println("Codility MaxProfit problem")
	a := []int {23171, 21011, 21123, 21366, 21013, 21367}
  	
  	var result int
    result = maxProfit(a)
    fmt.Println("The maximum possible profit from one transaction:" , result)
}


