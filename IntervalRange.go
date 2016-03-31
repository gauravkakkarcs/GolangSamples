/*
Write a function that takes a list of integer ranges and an additional range and efficiently prints all non-overlapping ranges from the additional range.
For example, given a list or ranges [6000000 – 10000000], [17000000 – 24000000], [15000000 – 19000000] and a new range [8000000 – 28000000], your function should print [10000001 – 14999999], [24000001 - 28000000].
The only guarantee is the range is valid i.e. range.start <= range.end.  Try to avoid making any other assumptions but if you do make other assumptions, make sure you mention those explicitly 
*/

package main

import "fmt"
import "sort"


type Range struct {
	first int
	second int
}
	
func out_range(ranges []Range, check Range) (out []Range) {
	
	out = make([] Range, 0, len(ranges)*2)
	
	var L  = make([]int, len(ranges)) 
	var R  = make([]int, len(ranges))

	for idx, val := range  ranges {
		L[idx],  R[idx] =  val.first, val.second
	}
	sort.Ints(L)
	sort.Ints(R)

	
	lc := 0
	rc := 0 
	ln := len(ranges)
	
	
	for  lc < (ln - 1) {
		
		if (R[rc] < L[lc+1]) {
			
			out = append(out, Range{R[rc]+1, L[lc+1]-1})
		}
		lc++ 
		rc++
	} 
	
	if(L[0] > check.first) {
		
    	out = append(out, Range{check.first, L[0] - 1})
    }
    
    if(R[len(R)-1] < check.second) {
    	
    	out = append(out, Range{R[len(R)-1] + 1, check.second})
    }
	
	return out	 
}

	
func main() {
	fmt.Println("Interval Range")
	
	ranges := []Range { 
		{60, 100}, 
		{170, 240}, 
		{150, 190},
	}
	
	check_range := Range {80, 280}
	
	var out []Range
	
	fmt.Printf("%d %d\n", ranges, len(ranges))

	out = out_range(ranges, check_range)
	
	fmt.Printf("%d\n", out)
}
