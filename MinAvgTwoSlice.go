// Solution to codility problem MinAvgTwoSlice: https://codility.com/programmers/task/min_avg_two_slice/
package main

import "fmt"



func minAvgTwoSlice(a []int) int {
	
	alen := len(a)
	if (alen <= 2){
		return 0;
	}
	
	var minavg float64 
	minavg = float64(a[0]+ a[1])/2.0 
		
	pos := 0
	for i := 0; i < (alen -2); i++  {
		 if( float64(a[i] + a[i+1])/2.0   < minavg) {
            minavg = float64(a[i] + a[i+1])/2.0
            pos = i
        }
        if( float64(a[i] + a[i+1] + a[i+2])/3.0  < minavg) {
            minavg = float64(a[i] + a[i+1] + a[i+2])/3.0
            pos = i
        }
	}
	
	if( float64(a[alen - 2] + a[alen - 1])/2.0   < minavg) {
        minavg = float64(a[alen - 2] + a[alen - 1])/2.0
        pos = alen- 2
    }
    
	
	return pos
}




func main() {
	fmt.Println("Codility MinAvgTwoSlice problem")
	a := []int {4, 2, 2, 5, 1, 5, 8}
  	
  	var result int
    result = minAvgTwoSlice(a)
    fmt.Println("The  starting position of a slice whose average is minimal is:" , result)
}
