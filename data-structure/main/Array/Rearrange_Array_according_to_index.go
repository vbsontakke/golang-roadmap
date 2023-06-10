package Array

// Given an array of elements of length N, ranging from 0 to N – 1. 
//All elements may not be present in the array. If the element is not present 
//then there will be -1 present in the array. 
//Rearrange the array such that A[i] = i and if i is not present, 
//display -1 at that place.

import "fmt"

func main(){
	input := []int{-1, -1, 6, 1, 9, 3, 2, -1, 4, -1}
	output := rearrange(input)
	fmt.Println("Input:",input)
	fmt.Println("Result:",output)
}
func rearrange(input []int)[]int{
	i:=0
	for i=0;i<len(input);i++{
		if input[i] != i && input[i] != -1{
			input[i], input[input[i]] = input[input[i]],input[i]
		}
	}
	return input
}