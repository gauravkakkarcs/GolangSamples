/*
Sketch out the code for a thread-safe queue, in particular the push_back and pop_front functions that would be suitable for a producer/consumer problem.   The consumer threads should be able to wait efficiently for items to be produced.
*/

package main

import "fmt"
import "sync"
import "time"

// Queue implementation
type Node struct {
	Value int
}

type Queue []*Node

func (q *Queue) Push(n *Node) {
    *q = append(*q, n)
}

func (q *Queue) Pop() (n *Node) {
    n = (*q)[0]
    *q = (*q)[1:]
    return
}

func (q *Queue) Len() int {
    return len(*q)
}

var q Queue


// Producer Consumer scenario
var cond *sync.Cond

func function_1() {
	fmt.Println("fn1")
	count := 10
	for count > 0 {
		cond.L.Lock()
		q.Push(&Node{count})
		cond.L.Unlock()
		cond.Signal()
		time.Sleep(time.Millisecond * 10)
		count--;
	}

}


func function_2() {
	fmt.Println("fn2")
	var data *Node
	data = &Node{Value: 0}
	 
	for data.Value != 1 {
		cond.L.Lock()
		cond.Wait()
		data = q.Pop()
		cond.L.Unlock()
		fmt.Printf("%d \n", data.Value);  
	}
}


func main() {
	fmt.Println("Producer Consumer")
	cond = &sync.Cond {L: &sync.Mutex{}}
	
	go function_1()
	go function_2()
	
	time.Sleep(time.Millisecond * 1000)
	fmt.Println("Over")
}
