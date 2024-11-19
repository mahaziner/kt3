// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var globVar int

// var Mut sync.Mutex

// func increment(WGroup *sync.WaitGroup) {
// 	defer WGroup.Done() 

// 	Mut.Lock()  
// 	globVar++  
// 	Mut.Unlock() 
// }

// func main() {
// 	var WGroup sync.WaitGroup

// 	for i := 0; i < 5; i++ {
// 		WGroup.Add(1) 
// 		go increment(&WGroup)
// 	}

// 	WGroup.Wait() 

// 	fmt.Printf("значение глобальной переменной: %d\n", globVar)
// }