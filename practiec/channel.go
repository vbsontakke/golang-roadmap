package main

import (
    "fmt"
    "sync"
)

func reciver(ch chan int, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("recieving...")
	
	for num := range ch {
		fmt.Println(num)
	}
}
func sender(ch chan int, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("sending...")
	for i := 0 ; i <10 ; i++ {
		ch <- i
	}
	close(ch)
	
}

func main(){
	ch := make(chan int,10)
	var wg sync.WaitGroup
	fmt.Println("starting...")
	wg.Add(1)
	go reciver(ch,&wg)
	wg.Add(1)
	go sender(ch,&wg)
	
	
	wg.Wait()
}