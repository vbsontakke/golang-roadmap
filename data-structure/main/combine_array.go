package main

import "fmt"

func main(){
	arr1 := []int{1,3,5,7,9}
	arr2 := []int{2,4,6,8}
	var result []int

	var counter1 int
	var counter2 int
	var counter3 int
	counter1,counter2,counter3=0
	for	(counter1<=len(arr1) || counter2<=len(arr2)){
		if(arr1[counter1]<arr2[counter2]){
			result[counter3]=arr1[counter1]
			counter1++
			counter3++
		}else if(arr1[counter1]>=arr2[counter2]){
			result[counter3]=arr2[counter2]
			counter2++
			counter3++
		}
	}
}