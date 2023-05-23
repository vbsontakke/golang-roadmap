package main

import (
	"fmt"
)

type node struct {
	data        int
	pointernext *node
}

func AppendNode(pointer *node, data int) *node {
	temp := node{data, nil}

	if pointer == nil {
		pointer = &temp
	} else {
		temppointer := pointer
		for temppointer.pointernext != nil {
			temppointer = temppointer.pointernext
		}
		temppointer.pointernext = &temp
	}
	return pointer
}

func printList(pointer *node) {
	temppointer := pointer
	for temppointer != nil {
		fmt.Println(temppointer.data)
		temppointer = temppointer.pointernext
	}
}

func findMiddle(pointer *node) int {
	temppointer := pointer
	temppointer2 := pointer
	for temppointer != nil && temppointer2.pointernext != nil {
		temppointer = temppointer.pointernext
		temppointer2 = temppointer2.pointernext.pointernext
	}
	return temppointer.data

}
func main() {
	var pointer = &(node{1, nil})

	pointer = AppendNode(pointer, 2)
	pointer = AppendNode(pointer, 3)
	pointer = AppendNode(pointer, 4)
	pointer = AppendNode(pointer, 5)
	pointer = AppendNode(pointer, 6)
	pointer = AppendNode(pointer, 7)
	// pointer = AppendNode(pointer,8)
	printList(pointer)

	middle := findMiddle(pointer)
	fmt.Println("Middle Element : ", middle)

}
