// Solution to codility problem Nesting: https://codility.com/programmers/task/nesting/
package main

import "fmt"



type stack []byte

func (s stack) Push(v byte) stack {
    return append(s, v)
}

func (s stack) Pop() (stack, byte) {
    l := len(s)
    return  s[:l-1], s[l-1]
}

func (s stack) Top() (byte) {
    l := len(s)
    return  s[l-1]
}


func nesting(a string) int {
	 s := make(stack, 0)
	 alen := len(a)
	 if alen == 0 {
	 	return 1
	 }
	 
	 for i := 0; i < alen; i++ {
	 	switch a[i] {
	 		case '(': 
	 			s = s.Push(a[i])
	 			break

	 		case ')':
	 			if (len(s) != 0) {
	 				ch := s.Top()
	 				 if((ch == '(') && a[i] == ')')  {
                        	s, _ = s.Pop()
                        } else {
                        	return 0;
                        }
	 			} else{
	 				return 0;
	 			}
	 			break;
	 	
	 		default: 
	 			return 0; 
	 	}
	 }
 	if len(s) == 0 {
 		return 1
 	} else {
 		return 0
 	}
}




func main() {
	fmt.Println("Codility Nesting problem")
	a := "(()(())())"
	b := "())" 
  	
  	var result int
    result = nesting(a)
    fmt.Println("The brackets", a, "are nested: " , result)
    result = nesting(b)
    fmt.Println("The brackets", b, "are nested:" , result)
    
}
