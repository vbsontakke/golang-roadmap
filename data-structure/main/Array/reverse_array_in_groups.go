package main

import(
	"fmt"
)

func main(){
	input := []int{1,2,3,4,5}
	k := 3
	n :=5
	result := reverse_in_groups(input,k,n)
	fmt.Println(result)
}

func reverse_in_groups(input []int,k int,n int)([]int){
	
	i,j := 0,k-1

	for i < j {
		input[i],input[j] = input[j],input[i]
		i++
		j--
	}

	return input
}