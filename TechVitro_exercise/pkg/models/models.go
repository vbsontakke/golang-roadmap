package models

// We will be declaring all the data types which we will be requiring while developing
// Project

// We will use Seat Allocation Struct for storing seat related data 
type seatAllocation struct{
	category string
	cost_per_ticket int
	list_of_seats map[string]bool
}

// We will use Show Allocation struct for storing the show related data
type showAllocation struct{
	showNumber int
	seatAllocationdetails seatAllocation
	serviceTax float32
	swachhBharatCess float32
	krishiKalyanCess float32
}

