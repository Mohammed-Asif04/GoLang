package main

import "fmt"

func main()  {

	// Strings

	var nameOne string = "mario";
	var nameTwo = "Asif"
	var namethree string

	nameOne = "Mohad"
	namethree = "Go lang"
	fmt.Println(nameOne ,nameTwo,namethree);

	nameFour := "learning Go" // only valid in func and another way of declaration of variable
	fmt.Println(nameFour);

	// Numbers

	var ageOne int = 20
	var ageTwo= 30
	ageThree :=40

	fmt.Println(ageOne,ageTwo,ageThree);

	//  Bits and Memeory

	// var numOne int8 = 25  // int(number) refers to all signed bits range 2^(number)/2 both plus and minus sign
    // var numTwo int8 = -128
	// var numThree uint8 = 250 // only take plus bits and remove minus bits range 

	// var scoreOne float32 =25.98
	// var scoreTwo float64= 2353253454363465645654.878
	
}