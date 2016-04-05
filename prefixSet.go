// Solution to codility problem PrefixSet: https://codility.com/programmers/task/prefix_set/
package main

import "fmt"



func prefixSet(a []int) int {
	
	index := 0
	alen := len(a)
	if (alen == 0){
		return 0;
	}
	if (alen == 1){
		return 0;
	}
	
	s := make(map[int]int)
	for i, v := range a {
		if _, ok := s[v]; !ok {
			s[v] = 1
			index = i
		}
	}

	return index
}




func main() {
	fmt.Println("Codility Prefix Set problem")
	a := []int {2, 2, 1, 0, 3, 1, 1, 1, 2, 1, 0, 3}
  	
  	var result int
    result = prefixSet(a)
    fmt.Println("The binary gap is:" , result)
}
