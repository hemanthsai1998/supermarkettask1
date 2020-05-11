package main

import (
	mart "cisco/supermarket"
	"fmt"
)

func main() {
	var a int32
	fmt.Println("Welcome to Supermarket")
	fmt.Println("\n1.Get price of Item\n2.Add Item\n3.Update value of Item\n4.Delete Item\n5.Show Items\n0.Exit")
Loop:
	for {
		var item string
		var value string
		fmt.Print("Enter your choice:")
		fmt.Scan(&a)
		switch a {
		case 1:
			fmt.Print("Item name?:")
			fmt.Scan(&item)
			mart.Get(item)
		case 2:
			fmt.Print("Item name and value?:")
			fmt.Scanf("%s %s", &item, &value)
			mart.Post(item, value)
		case 3:
			fmt.Print("Item name and updated value?:")
			fmt.Scanf("%s %s", &item, &value)
			mart.Put(item, value)
		case 4:
			fmt.Print("Item name?:")
			fmt.Scan(&item)
			mart.Delete(item)
		case 5:
			mart.Print()
		case 0:
			fmt.Println("exiting...")
			break Loop
		default:
			fmt.Println("not valid..select from 0 to 5")
		}
	}
}
