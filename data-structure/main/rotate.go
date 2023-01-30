package main

import "fmt"

func rotate(input []int)[]int{
	var temp int
	var result[]int
	result = input
	temp = input[0]
	for i,_:= range input{
		if (i == len(input)-1){
			break
		}
		result[i]=input[i+1]
	}
	result[len(input)-1] = temp
	
	return result
}
func main(){
	input := []int{1,2,3,4,5,6,7,8}
	input = rotate(input)
	fmt.Println(rotate(input))

}
