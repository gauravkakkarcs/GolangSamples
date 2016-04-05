// Solution to codility problem EquiLeader: https://codility.com/programmers/task/equi_leader/
package main

import "fmt"




func equiLeader(a []int) int {
	
	alen := len(a)
	if (alen == 0){
		return -1;
	}
	if (alen == 1){
		return 0;
	}
	count, index  := 0, 0
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
    
    total := 0;
   	for index = 0; index < alen; index++ {
       if (maj == a[index]) {
       total++;
       }
   }
   
   if(total <= (alen >> 1)) {
   	return 0;
   	}
   
   L, R, equi := 0, 0, 0
   for index = 0; index < alen; index++ {
        if(a[index] == maj) {
            L++;
        }
        R = total - L;
        if((L > (index+1)/2)  && (R > (alen-index-1)/2)) {
            equi++;
        }
   }
    
    return equi
	
}




func main() {
	fmt.Println("Codility EquiLeader problem")
	a := []int {4, 3, 4, 4, 4, 2}
  	
  	var result int
    result = equiLeader(a)
    fmt.Println("The number of equi leaders are:" , result)
}


