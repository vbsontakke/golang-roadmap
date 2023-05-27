package main

import (
	"fmt"
)

type Area interface{
	CalculateArea()int
}

type Square struct{
	length int
}

type Rectangle struct{
	length, breadth int
}

func (square Square)Area(length int)int{
	return length*length
}

func (rectangle Rectangle)Area(length int,breadth int)int{
	return length*breadth
}

func main(){
	square := Square{23}
	areaofSquare := square.Area(square.length)
	fmt.Println("Area of square with Length",square.length,"is ",areaofSquare)

	rectangel := Rectangle{23,20}
	areaofRectangle := rectangel.Area(rectangel.length,rectangel.breadth)
	fmt.Println("Area of square with Length",rectangel.length, "and Breadth ",rectangel.breadth,"is ",areaofRectangle)


}