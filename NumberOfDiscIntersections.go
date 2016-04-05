// Solution to codility problem NumberOfDiscIntersections: https://codility.com/programmers/task/number_of_disc_intersections/
package main

import "fmt"
import "sort"



func numberOfDiscIntersections(a []int) int {
	
	alen := len(a)
	if (alen < 2){
		return 0;
	}
	var (
		overlaps int64
		total int64
	)
	overlaps, total = 0, 0
	l := make([]int, alen)
	r := make([]int, alen)
	
	for i, v := range a {
		l[i] = i - v;
		r[i] = i + v;
	}
	sort.Ints(l)
	sort.Ints(r)
	
	lc, rc := 0, 0 
	for(lc < alen && rc < alen) {
		
		if(l[lc] <= r[rc]){
            overlaps += total
            total++
            lc++
        } else {
            total--
            rc++
        }
	}
	
	return int(overlaps)
}




func main() {
	fmt.Println("Codility NumberOfDiscIntersections problem")
	a := []int {1, 5, 2, 1, 4, 0}
  	
  	var result int
    result = numberOfDiscIntersections(a)
    fmt.Println("The number of (unordered) pairs of intersecting discs is: " , result)
}
