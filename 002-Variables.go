package main

import (
	"fmt"
	"strconv"
)


//  variables declared outside the main function are package level variables 
// Lower case package level variablescan be accesed 
// by any file in that package
// var packageLevelVar int = 71 

// upper case package lvl vars are exported from that package and can be accessed
// globally
// var GlobalVar int = 93

func main() {
	// There are three methods to declare a variable in GO
	// 1> Variable name followed by "var" keyword followed by type
	// 2> Step 1 but also assigning value 
	// 3 > Using := operator 
	
	var a int = 23
	var b float32 
	b = 17.3
	c := 29
	fmt.Println(a, b, c)

	// formating output
	fmt.Printf("%v, %T \n", b, b)

	// You could also wrap similar variables in a block
	var(
		name string = "Arjun"
		age int = 21
		height float32 = 5.11
	)

	fmt.Println(name, age, height)


	// Converting one var type to another

	var intVar int = 7
	var floatConverted float32;
	fmt.Printf("%v, %T, \n", intVar, intVar)

	floatConverted = float32(intVar)
	fmt.Printf("%v, %T \n", floatConverted, floatConverted)

	var strVar string = strconv.Itoa(intVar)
	fmt.Printf("%v, %T \n", strVar, strVar)


}