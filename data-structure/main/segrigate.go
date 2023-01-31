package main

import "fmt"

func segrigate(input []int)[]int{
	start :=0
	end := len(input)-1
	for (start<end){
		if (input[start]==1 && input[end]==0){
			input[start]=0
			input[end]=1
			start++
			end--
		}else if(input[start]==0){

			start++
		}else if(input[end]==1){
			end--
		}
	}
	return input
}

func main(){
	input := []int{1,0,0,0,1,1,1}
	
	fmt.Println(segrigate(input))
}