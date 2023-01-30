package main

import "fmt"

func main(){
	input := []int{1,2,3,4,5,6,7,8}

	for i,v := range input{
		fmt.Println(v)
	}
}