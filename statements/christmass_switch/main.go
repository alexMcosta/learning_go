package main

import "fmt"

func main() {

	fmt.Println("On the 1st day of Christmas my true love gave to me: \n A Partridge in a pear tree \n \n")
	//basic switch statement
	for day := 2; day <= 12; day++ {

		fmt.Println("On the", day, "day of Christmass my true love gave to me:")

		switch day {
		case 12:
			fmt.Println("Twelve Drummers Drumming")
			fallthrough
		case 11:
			fmt.Println("Eleven Pipers Piping")
			fallthrough
		case 10:
			fmt.Println("Ten Lords a Leaping")
			fallthrough
		case 9:
			fmt.Println("Nine Ladies Dancing")
			fallthrough
		case 8:
			fmt.Println("Eight Maids a Milking")
			fallthrough
		case 7:
			fmt.Println("Seven Swans a Swimming")
			fallthrough
		case 6:
			fmt.Println("Six Geese a Laying")
			fallthrough
		case 5:
			fmt.Println("Five Golden Rings")
			fallthrough
		case 4:
			fmt.Println("Four Calling Birds")
			fallthrough
		case 3:
			fmt.Println("Three French Hens")
			fallthrough
		case 2:
			fmt.Println("Two Turtle Doves")
			fallthrough
		case 1:
			fmt.Println("and a Partrifge in a Pear Tree \n \n")
		}
	}
}
