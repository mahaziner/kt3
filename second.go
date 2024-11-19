// package main

// import (
// 	"fmt"
// 	"strconv"
// 	"sync"
// )

// func main() {

// 	nummyChannel := make(chan int)    
// 	strmyChannel := make(chan string) 


// 	var WGroup sync.WaitGroup

// 	go func() {
// 		for i := 1; i <= 10; i++ {
// 			nummyChannel <- i
// 		}
// 		close(nummyChannel) 
// 	}()


// 	for i := 0; i < 10; i++ {
// 		WGroup.Add(1) 
// 		go func() {
// 			defer WGroup.Done() 
// 			for num := range nummyChannel {
// 				strmyChannel <- strconv.Itoa(num)
// 			}
// 		}()
// 	}

// 	go func() {
// 		WGroup.Wait()
// 		close(strmyChannel)
// 	}()


// 	for str := range strmyChannel {
// 		fmt.Println(str)
// 	}
// }
