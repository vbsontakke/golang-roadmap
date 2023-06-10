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

func detectloop(head *node)bool{
	slowpointer := head
	fastpointer := head.pointernext.pointernext
	for slowpointer != nil && fastpointer.pointernext != nil {
		if slowpointer == fastpointer{
			return true
		}
		slowpointer = slowpointer.pointernext
		fastpointer = fastpointer.pointernext.pointernext
	}
	return false
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

	isLoop := detectloop(pointer)
	if isLoop{
		fmt.Println("Loop Detected ")
	}else{
		fmt.Println("Loop Not Detected ")
	}
	

}