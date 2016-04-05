// Solution to codility problem BinaryGap: https://codility.com/programmers/task/binary_gap/

package main

import "fmt"



func binaryGap(n int) int {
	
	if (n < 5){
		return 0
	}

	// enum way in golang:
	type state int
	const (
		first_0 state  = iota
		first_1
		gap_started_0
	)
	
	var num int
	num = n 
	max_gap, cur_gap := 0, 0
	var ( 
		cur_state state
		next_state state
	)
	
	if(num & 0x01 == 1) {
		cur_state = first_1
	} else {
    	cur_state = first_0
    }
    
    num = num >> 1
    for ; num != 0; num >>= 1 {
    	switch cur_state {
    		case first_0:
    		 	if(num & 0x01 == 1) {
                    next_state = first_1
                } else {
                    next_state = first_0
                }                
                break
                
    		case first_1:
    			if(num & 0x01 == 1) {
                    next_state = first_1
                } else {
                    cur_gap = 1
                    next_state = gap_started_0
                }                
                break
                
    		case gap_started_0:
    			if(num & 0x01 == 1) {
                    next_state = first_1
                    if(max_gap < cur_gap) {
                    	max_gap = cur_gap
                    }
                }  else {
                    cur_gap++
                    next_state = gap_started_0
                }
                 break
    	}
    	 cur_state = next_state
    }
	
	

	return max_gap
}




func main() {
	fmt.Println("Codility Binary Gap problem")
	a := 1041
  	
  	var result int
    result = binaryGap(a)
    fmt.Println("The binary gap is:" , result)
}
