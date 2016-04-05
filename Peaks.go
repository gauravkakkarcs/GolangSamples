// Solution to codility problem Peaks: https://codility.com/programmers/task/peaks/
package main

import "fmt"


func peaks(a []int) int {
	alen := len(a)
	if (alen < 3){
		return 0
	}
	
	peaks := make([]int, 0)
    
    for i := 1; i < alen - 1; i++ {
      if(a[i] > a[i-1] && a[i] > a[i+1]) {
      	
          peaks = append(peaks, i)
      }
    }

    for size := 1; size <= alen; size++ {
        if(alen % size != 0) {
        	continue // skip if non-divisible
        }
        
        find := 0
        groups := alen/size
        ok := true
        
        // Find whether every group has a peak
        for peakidx := 0; peakidx < len(peaks); peakidx++ {
        	
            // index stored (if not present in current block -- find --> break
            if(peaks[peakidx]/size > find) {
                ok = false
                break
            }
            
            // if present, incr find for next iteration.
            if(peaks[peakidx]/size == find) {
                find++
            }
        }
        
        // if no of blocks find is having all blocks, then return out of it. 
        if(find != groups) {   
            ok = false
        }
        if(ok) {
            return groups
        }
    }
    return 0
}




func main() {
	fmt.Println("Codility Peaks problem")
	a := []int {1,2,3,4,3,4,1,2,3,4,6,2}
  	
  	var result int
    result = peaks(a)
    fmt.Println("The maximum number of blocks that array can be divided into  is:" , result)
}


