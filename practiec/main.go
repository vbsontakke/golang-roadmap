package main

import(
	"fmt"
)

type student struct{
	name string
	roll_no int
	class int
}

type class_student interface{
	get_roll_no(name string)(int)
}

func (student1 student ) get_roll_no()(roll_no int){
	return student1.roll_no
}

func main(){
	student1 := student{"vinayak",1,12}
	roll_no := student1.get_roll_no()
	fmt.Println("roll call :",roll_no)
}