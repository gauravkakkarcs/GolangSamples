/*
1.) Typical fibonacci sequence using GO channels
2.) Routine to remove dup elements from the string array
 */

package main

import "fmt"


func fibonacci(c, quit chan int) {
    x, y := 0, 1
    for {
        select {
        case c <- x:
            x, y = y, x+y
        case <-quit:
            fmt.Println("quit")
            return
        }
    }
}


func UniqStr(col []string) []string {
	m := map[string]struct{}{}
	for _, v := range col {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
		}
	}
	list := make([]string, len(m))
	fmt.Println(m);
	
	i := 0
	for v := range m {
		list[i] = v
		i++
	}
	return list
}


func main() {

	
// 1

    ch := make(chan int)
    quit := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(<-ch)
        }
        quit <- 0
    }()
    fibonacci(ch, quit)

// 2
	

    data:= []string{"a", "b", "c", "a", "b", "c"}
    fmt.Println(UniqStr(data));

}
