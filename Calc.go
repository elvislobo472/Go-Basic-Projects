package main

import "fmt"




func add (x, y int) int{
	return x + y
}


func sub (x, y int) int{
	return x - y
}


func multiply (x, y int) int{
	return x * y
}

func divide (x, y int) int{
	return x / y
}

func modulus (x, y int) int{
	return x % y
}




func main(){


	fmt.Println("===============================Welcome to Basic Calculator in GO===============================")






	var num1, num2 int

	fmt.Print("Enter Number 1 and Number 2: ")
	fmt.Scan(&num1, &num2)


	// fmt.Println(num1, num2)
	var choice int
	fmt.Print("Enter your choice (1-5):")

	fmt.Scan(&choice)


	switch choice{
	case 1:
		fmt.Print("Result of addition is ",add(num1, num2))
	
	case 2:
		fmt.Print("Result of subtraction is ", sub(num1, num2))
	

	case 3:
		fmt.Print("Result of multiplication is ", multiply(num1, num2))

	case 4:
		fmt.Print("Result of division is ", divide(num1, num2))

	case 5:
		fmt.Print("Result of Modulus is ", modulus(num1, num2))

	default:
		fmt.Print("Please a enter a valid choice.")
	


	}



}